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
func Index(mstr *template.Template, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		tkn := r.Context().Value("token")

		if tkn == nil {
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

		result := struct {
			Token string
			CMS   map[string]interface{}
		}{
			CMS:   content,
			Token: tkn.(string),
		}

		err = ctx.Serve(http.StatusOK, mix.Page("index", result, ctx.GetTokenInfo(), mstr, tmpl))

		if err != nil {
			log.Println(err)
		}
	}
}
