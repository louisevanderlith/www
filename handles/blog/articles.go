package blog

import (
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/www/resources"
	"html/template"
	"log"
	"net/http"
)

func GetArticles(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Articles", "./views/articles.html")

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)
		src := resources.APIResource(http.DefaultClient, ctx)

		result, err := src.FetchArticles("A10")

		if err != nil {
			log.Println("Fetch Articles Error", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		err = ctx.Serve(http.StatusOK, pge.Page(result, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchArticles(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Articles", "./views/articles.html")

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)
		src := resources.APIResource(http.DefaultClient, ctx)

		result, err := src.FetchArticles(ctx.FindParam("pagesize"))

		if err != nil {
			log.Println("Fetch Articles Error", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		err = ctx.Serve(http.StatusOK, pge.Page(result, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewArticle(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Articles View", "./views/articlesView.html")

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)
		key, err := husk.ParseKey(ctx.FindParam("key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, ctx)

		result, err := src.FetchArticle(key.String())

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = ctx.Serve(http.StatusOK, pge.Page(result, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println(err)
		}

		//TODO: comments
	}
}
