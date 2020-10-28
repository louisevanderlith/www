package handles

import (
	"github.com/louisevanderlith/blog/api"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"html/template"
	"log"
	"net/http"
)

func GetArticles(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Articles", tmpl, "./views/articles.html")

	return func(w http.ResponseWriter, r *http.Request) {
		result, err := api.FetchLatestArticles(CredConfig.Client(r.Context()), Endpoints["blog"], "A10")

		if err != nil {
			log.Println("Fetch Articles Error", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchArticles(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Articles", tmpl, "./views/articles.html")

	return func(w http.ResponseWriter, r *http.Request) {
		result, err := api.FetchLatestArticles(CredConfig.Client(r.Context()), Endpoints["blog"], drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println("Fetch Articles Error", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewArticle(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Articles View", tmpl, "./views/articlesView.html")

	return func(w http.ResponseWriter, r *http.Request) {
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		result, err := api.FetchArticle(CredConfig.Client(r.Context()), Endpoints["blog"], key)

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println(err)
		}

		//TODO: comments
	}
}
