package handles

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/coreos/go-oidc"
	"github.com/louisevanderlith/blog/api"
	"github.com/louisevanderlith/blog/core"
	commentapi "github.com/louisevanderlith/comment/api"
	"github.com/louisevanderlith/comment/core/commenttype"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/husk/records"
	"golang.org/x/oauth2"
	"io/ioutil"
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

		err = mix.Write(w, fact.Create(r, "Articles", "./views/articles.html", mix.NewDataBag(data)))

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

		err = mix.Write(w, fact.Create(r, "Articles", "./views/articles.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func FetchArticleWithSource(host string, k hsk.Key, tokenSource oauth2.TokenSource) (core.Article, error) {
	url := fmt.Sprintf("%s/articles/%s", host, k.String())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return core.Article{}, fmt.Errorf("oidc: create GET request: %v", err)
	}

	token, err := tokenSource.Token()
	if err != nil {
		return core.Article{}, fmt.Errorf("oidc: get access token: %v", err)
	}

	token.SetAuthHeader(req)

	resp, err := http.DefaultClient.Do(req) //doRequest(ctx, req)

	if err != nil {
		return core.Article{}, err
	}
	//resp, err := web.Get(url)

	//if err != nil {
	//	return core.Article{}, err
	//}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return core.Article{}, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := core.Article{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func ViewArticle(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		tkn := r.Context().Value("Token")
		//var tknsrc oauth2.TokenSource
		var clnt *http.Client

		if tkn != nil {
			authTkn := tkn.(oauth2.Token)
			clnt = AuthConfig.Client(r.Context(), &authTkn)
			//tknsrc = AuthConfig.TokenSource(r.Context(), &authTkn)
		} else {
			clnt = credConfig.Client(r.Context())
			//tknsrc = credConfig.TokenSource(r.Context())
		}

		//clnt := credConfig.Client(r.Context())
		article, err := api.FetchArticle(clnt, Endpoints["blog"], key)

		//article, err := FetchArticleWithSource(Endpoints["blog"], key, tknsrc)
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

		//pge.ChangeTitle("Article " + article.Title)
		err = mix.Write(w, fact.Create(r, "Article View", "./views/articleview.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func getUserGravatar(r *http.Request) string {
	tknVal := r.Context().Value("IDToken")

	if tknVal == nil {
		log.Println("ID Token not set")
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
