package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/menu"
	"github.com/louisevanderlith/kong/middle"
	"github.com/louisevanderlith/www/handles/blog"
	"net/http"
)

func FullMenu(sectionAHead, sectionBHead, infoHead string) *menu.Menu {
	m := menu.NewMenu()

	m.AddItem(menu.NewItem("a", "#home", "Home", nil))
	m.AddItem(menu.NewItem("b", "#sectionA", sectionAHead, nil))
	m.AddItem(menu.NewItem("c", "#sectionB", sectionBHead, nil))
	m.AddItem(menu.NewItem("d", "#info", infoHead, nil))
	m.AddItem(menu.NewItem("e", "#services", "Services we Offer", nil))
	m.AddItem(menu.NewItem("e", "/blog", "Blog", nil))
	m.AddItem(menu.NewItem("f", "#contact", "Contact Us", nil))

	return m
}

func SetupRoutes(clnt, scrt, securityUrl, managerUrl, authorityUrl string) http.Handler {
	tmpl, err := drx.LoadTemplate("./views")

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	distPath := http.FileSystem(http.Dir("dist/"))
	fs := http.FileServer(distPath)
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", fs))

	clntIns := middle.NewClientInspector(clnt, scrt, http.DefaultClient, securityUrl, managerUrl, authorityUrl)
	r.HandleFunc("/", clntIns.Middleware(Index(tmpl), map[string]bool{"cms.content.view": true, "stock.services.search": true})).Methods(http.MethodGet)

	r.HandleFunc("/blog", clntIns.Middleware(blog.GetArticles(tmpl), map[string]bool{"blog.articles.search": true})).Methods(http.MethodGet)
	r.HandleFunc("/blog/{pagesize:[A-Z][0-9]+}", clntIns.Middleware(blog.SearchArticles(tmpl), map[string]bool{"blog.articles.search": true})).Methods(http.MethodGet)
	r.HandleFunc("/blog/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", clntIns.Middleware(blog.SearchArticles(tmpl), map[string]bool{"blog.articles.search": true})).Methods(http.MethodGet)
	r.HandleFunc("/blog/{key:[0-9]+\\x60[0-9]+}", clntIns.Middleware(blog.ViewArticle(tmpl), map[string]bool{"blog.articles.view": true})).Methods(http.MethodGet)

	return r
}
