package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/www/resources"
	"html/template"
	"log"
	"net/http"
)

//GetDefault returns the 'defaultsite'
func Index(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Index", tmpl, "./views/index.html")

	return func(w http.ResponseWriter, r *http.Request) {
		tkn := drx.GetToken(r)

		if len(tkn) == 0 {
			http.Error(w, "no token", http.StatusUnauthorized)
			return
		}

		src := resources.APIResource(http.DefaultClient, r)
		content, err := src.FetchProfileDisplay()

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		services, err := src.FetchServices("A6")

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		log.Println("services:", services)
		content["Services"] = services
		sectA := content["SectionA"].(map[string]interface{})
		sectB := content["SectionB"].(map[string]interface{})
		info := content["Info"].(map[string]interface{})

		tknInfo := drx.GetIdentity(r)
		pge.ChangeTitle(tknInfo.GetProfile())
		pge.AddMenu(FullMenu(sectA["Heading"].(string), sectB["Heading"].(string), info["Heading"].(string)))

		err = mix.Write(w, pge.Create(r, content))

		if err != nil {
			log.Println(err)
		}
	}
}
