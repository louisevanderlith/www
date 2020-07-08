package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/menu"
	"github.com/louisevanderlith/kong"
	"github.com/louisevanderlith/www/handles/blog"
	"net/http"
)

func FullMenu(sectionAHead, sectionBHead, infoHead string) *menu.Menu {
	m := menu.NewMenu()

	m.AddItem(menu.NewItem("a", "#home", "Home", nil))
	m.AddItem(menu.NewItem("b", "#sectionA", sectionAHead, nil))
	m.AddItem(menu.NewItem("c", "#sectionB", sectionBHead, nil))
	m.AddItem(menu.NewItem("d", "#info", infoHead, nil))
	m.AddItem(menu.NewItem("e", "#services", "What we Offer", nil))
	//m.AddItem(menu.NewItem("e", "/blog", "Blog", nil))
	m.AddItem(menu.NewItem("f", "#contact", "Contact Us", nil))

	return m
}

func SetupRoutes(clnt, scrt, secureUrl string) http.Handler {
	tmpl, err := droxolite.LoadTemplate("./views")

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	distPath := http.FileSystem(http.Dir("dist/"))
	fs := http.FileServer(distPath)
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", fs))

	r.HandleFunc("/", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, "", Index(tmpl), "cms.content.view", "stock.services.search")).Methods(http.MethodGet)

	r.HandleFunc("/blog", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, "", blog.GetArticles(tmpl), "blog.articles.search")).Methods(http.MethodGet)
	r.HandleFunc("/blog/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, "", blog.SearchArticles(tmpl), "blog.articles.search")).Methods(http.MethodGet)
	r.HandleFunc("/blog/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, "", blog.SearchArticles(tmpl), "blog.articles.search")).Methods(http.MethodGet)
	r.HandleFunc("/blog/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, "", blog.ViewArticle(tmpl), "blog.articles.view")).Methods(http.MethodGet)

	return r
}
