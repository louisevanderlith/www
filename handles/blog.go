package handles

import (
	"github.com/louisevanderlith/blog/api"
	"github.com/louisevanderlith/blog/core"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"html/template"
	"log"
	"net/http"
	"time"
)

func GetArticles(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Articles", tmpl, "./views/articles.html")
	//pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := api.FetchLatestArticles(credConfig.Client(r.Context()), Endpoints["blog"], "A10")

		if err != nil {
			log.Println("Fetch Articles Error", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		pge.ChangeTitle("Blog")
		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchArticles(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Articles", tmpl, "./views/articles.html")

	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())

	return func(w http.ResponseWriter, r *http.Request) {
		pgSize := drx.FindParam(r, "pagesize")
		result, err := api.FetchLatestArticles(credConfig.Client(r.Context()), Endpoints["blog"], pgSize)

		if err != nil {
			log.Println("Fetch Articles Error", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		pge.ChangeTitle("Blog")
		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewArticle(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Articles View", tmpl, "./views/articleview.html")
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		article, err := api.FetchArticle(credConfig.Client(r.Context()), Endpoints["blog"], key)

		if err != nil {
			log.Println("Fetch Article Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		result := struct {
			CreateDate time.Time
			Article    core.Article
			Comments   []interface{}
		}{
			CreateDate: key.GetTimestamp(),
			Article:    article,
		}

		pge.ChangeTitle(article.Title)
		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}

		//TODO: comments
	}
}
