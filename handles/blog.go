package handles

import (
	"crypto/md5"
	"fmt"
	"github.com/coreos/go-oidc"
	"github.com/louisevanderlith/blog/api"
	"github.com/louisevanderlith/blog/core"
	commentapi "github.com/louisevanderlith/comment/api"
	"github.com/louisevanderlith/comment/core/commenttype"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/husk/records"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetArticles(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := api.FetchLatestArticles(credConfig.Client(r.Context()), Endpoints["blog"], "A10")

		if err != nil {
			log.Println("Fetch Articles Error", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		bag := mix.NewDataBag(data)
		bag.SetValue("Title", "Articles")
		err = mix.Write(w, fact.Create(r, "Articles", "./views/articles.html", bag))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchArticles(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pgSize := drx.FindParam(r, "pagesize")
		data, err := api.FetchLatestArticles(credConfig.Client(r.Context()), Endpoints["blog"], pgSize)

		if err != nil {
			log.Println("Fetch Articles Error", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		bag := mix.NewDataBag(data)
		bag.SetValue("Title", "Articles")
		err = mix.Write(w, fact.Create(r, "Articles", "./views/articles.html", bag))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewArticle(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)

		article, err := api.FetchArticle(clnt, Endpoints["blog"], key)

		if err != nil {
			log.Println("Fetch Article Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		comments, err := commentapi.FetchCommentsFor(clnt, Endpoints["comment"], commenttype.Article, key)

		if err != nil {
			log.Println("Fetch Comments Error", err)
		}

		data := struct {
			CreateDate time.Time
			Article    core.Article
			Comments   records.Page
			Gravatar   string
		}{
			CreateDate: key.GetTimestamp(),
			Article:    article,
			Comments:   comments,
			Gravatar:   getUserGravatar(r),
		}

		bag := mix.NewDataBag(data)
		bag.SetValue("Title", "Article "+article.Title)
		err = mix.Write(w, fact.Create(r, "Article View", "./views/articleview.html", bag))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func getUserGravatar(r *http.Request) string {
	tknVal := r.Context().Value("IDToken")

	if tknVal == nil {
		return ""
	}

	idToken := tknVal.(*oidc.IDToken)
	claims := make(map[string]interface{})
	err := idToken.Claims(&claims)

	if err != nil {
		log.Println("Claims Error", err)
		return ""
	}

	email := claims["email"].(string)

	if len(email) == 0 {
		return ""
	}

	gravatar := md5.Sum([]byte(strings.ToLower(strings.Trim(email, " "))))

	return fmt.Sprintf("%x", gravatar)
}
