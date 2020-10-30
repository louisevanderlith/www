package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	folio "github.com/louisevanderlith/folio/api"
	"github.com/louisevanderlith/folio/core"
	"github.com/louisevanderlith/husk/records"
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
	return func(w http.ResponseWriter, r *http.Request) {
		clnt := CredConfig.Client(r.Context())

		result := struct {
			Content  core.Content
			Services records.Page
		}{}
		content, err := folio.FetchDisplay(clnt, Endpoints["folio"])

		if err != nil {
			log.Println("Fetch Profile Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		result.Content = content

		services, err := stock.FetchAllServices(clnt, Endpoints["stock"], "A6")

		if err != nil {
			log.Println("Fetch Services Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		result.Services = services

		sectA := content.SectionA
		sectB := content.SectionB
		info := content.Info

		pge.AddMenu(FullMenu(sectA.Heading, sectB.Heading, info.Heading))

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
