package handles

import (
	"github.com/coreos/go-oidc"
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/menu"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/open"
	folio "github.com/louisevanderlith/folio/api"
	"github.com/louisevanderlith/theme/api"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"log"
	"net/http"
)

var (
	AuthConfig *oauth2.Config
	credConfig *clientcredentials.Config
	Endpoints  map[string]string
)

func FullMenu(sectionAHead, sectionBHead, infoHead string) *menu.Menu {
	m := menu.NewMenu()

	m.AddItem(menu.NewItem("b", "/#sectionA", sectionAHead, nil))
	m.AddItem(menu.NewItem("c", "/#sectionB", sectionBHead, nil))
	m.AddItem(menu.NewItem("d", "/#info", infoHead, nil))
	m.AddItem(menu.NewItem("e", "/#services", "Products we Offer", nil))
	m.AddItem(menu.NewItem("e", "/blog", "Blog", nil))
	m.AddItem(menu.NewItem("f", "/#contact", "Contact Us", nil))

	return m
}

func SetupRoutes(host, clientId, clientSecret string, endpoints map[string]string) http.Handler {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, endpoints["issuer"])

	if err != nil {
		panic(err)
	}

	Endpoints = endpoints

	AuthConfig = &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  host + "/callback",
		Scopes:       []string{oidc.ScopeOpenID, "blog-view", "comment-save"},
	}

	credConfig = &clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     provider.Endpoint().TokenURL,
		Scopes:       []string{oidc.ScopeOpenID, "upload-artifact", "theme", "blog-view", "stock-view"},
	}

	err = api.UpdateTemplate(credConfig.Client(ctx), endpoints["theme"])

	if err != nil {
		panic(err)
	}

	tmpl, err := drx.LoadTemplate("./views")

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	distPath := http.FileSystem(http.Dir("dist/"))
	fs := http.FileServer(distPath)
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", fs))

	lock := open.NewUILock(provider, AuthConfig)

	r.HandleFunc("/login", lock.Login).Methods(http.MethodGet)
	r.HandleFunc("/callback", lock.Callback).Methods(http.MethodGet)
	gmw := open.NewGhostware(credConfig)

	fact := mix.NewPageFactory(tmpl)
	fact.AddModifier(mix.EndpointMod(Endpoints))
	fact.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	fact.AddModifier(ThemeContentMod())

	r.HandleFunc("/", gmw.GhostMiddleware(Index(fact))).Methods(http.MethodGet)

	r.HandleFunc("/blog", gmw.GhostMiddleware(GetArticles(fact))).Methods(http.MethodGet)
	r.HandleFunc("/blog/{pagesize:[A-Z][0-9]+}", gmw.GhostMiddleware(SearchArticles(fact))).Methods(http.MethodGet)
	r.HandleFunc("/blog/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", gmw.GhostMiddleware(SearchArticles(fact))).Methods(http.MethodGet)
	r.Handle("/blog/{key:[0-9]+\\x60[0-9]+}", lock.NoLoginMiddleware(ViewArticle(fact))).Methods(http.MethodGet)

	return r
}

func ThemeContentMod() mix.ModFunc {
	return func(b mix.Bag, r *http.Request) {
		clnt := credConfig.Client(r.Context())

		content, err := folio.FetchDisplay(clnt, Endpoints["folio"])

		if err != nil {
			log.Println("Fetch Profile Error", err)
			panic(err)
			return
		}

		b.SetValue("Folio", content)
		b.SetValue("Menu", FullMenu(content.SectionA.Heading, content.SectionB.Heading, content.Info.Heading))
	}
}
