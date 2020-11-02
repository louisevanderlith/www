package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	stock "github.com/louisevanderlith/stock/api"
	"html/template"
	"log"
	"net/http"
)

//GetDefault returns the 'defaultsite'
func Index(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Index", tmpl, "./views/index.html")
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		clnt := CredConfig.Client(r.Context())
		services, err := stock.FetchAllServices(clnt, Endpoints["stock"], "A6")

		if err != nil {
			log.Println("Fetch Services Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = mix.Write(w, pge.Create(r, services))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
