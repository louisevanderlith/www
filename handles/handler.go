package handles

import (
	"github.com/coreos/go-oidc"
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/menu"
	"github.com/louisevanderlith/droxolite/open"
	"github.com/louisevanderlith/theme/api"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

var (
	CredConfig *clientcredentials.Config
	Endpoints  map[string]string
	//FolioURL   string
	//BlogURL    string
	//StockURL   string
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

func SetupRoutes(host, clientId, clientSecret string, endpoints map[string]string) http.Handler {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, endpoints["issuer"])

	if err != nil {
		panic(err)
	}

	Endpoints = endpoints

	authConfig := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  host + "/callback",
		Scopes:       []string{oidc.ScopeOpenID},
	}

	CredConfig = &clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     provider.Endpoint().TokenURL,
		Scopes:       []string{oidc.ScopeOpenID, "artifact", "theme"},
	}

	err = api.UpdateTemplate(CredConfig.Client(ctx), endpoints["theme"])

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

	lock := open.NewUILock(authConfig)
	r.HandleFunc("/login", lock.Login).Methods(http.MethodGet)
	r.HandleFunc("/callback", lock.Callback).Methods(http.MethodGet)

	r.HandleFunc("/", GhostMiddleware(Index(tmpl))).Methods(http.MethodGet)

	r.HandleFunc("/blog", GhostMiddleware(GetArticles(tmpl))).Methods(http.MethodGet)
	r.HandleFunc("/blog/{pagesize:[A-Z][0-9]+}", GhostMiddleware(SearchArticles(tmpl))).Methods(http.MethodGet)
	r.HandleFunc("/blog/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", GhostMiddleware(SearchArticles(tmpl))).Methods(http.MethodGet)
	r.HandleFunc("/blog/{key:[0-9]+\\x60[0-9]+}", GhostMiddleware(ViewArticle(tmpl))).Methods(http.MethodGet)

	return r
}

func GhostMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tkn, err := CredConfig.Token(r.Context())

		if err != nil {
			panic(err)
		}

		acc := context.WithValue(r.Context(), "Token", tkn.AccessToken)
		next.ServeHTTP(w, r.WithContext(acc))
	}
}
