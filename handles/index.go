package handles

import (
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/www/resources"
	"html/template"
	"log"
	"net/http"
)

//GetDefault returns the 'defaultsite'
func Index(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Index", "./views/index.html")

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		tkn := ctx.GetToken()

		if len(tkn) == 0 {
			http.Error(w, "no token", http.StatusUnauthorized)
			return
		}

		src := resources.APIResource(http.DefaultClient, ctx)
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

		tknInfo := ctx.GetTokenInfo()
		pge.ChangeTitle(tknInfo.GetProfile())
		pge.AddMenu(FullMenu(sectA["Heading"].(string), sectB["Heading"].(string), info["Heading"].(string)))

		err = ctx.Serve(http.StatusOK, pge.Page(content, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println(err)
		}
	}
}
