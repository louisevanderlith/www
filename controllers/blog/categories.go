package blog

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/husk"
)

type Categories struct {
}

func (req *Categories) Get(ctx context.Requester) (int, interface{}) {
	//Show categories...
	//req.Setup("categories", "Blog Categories", false)
	return http.StatusOK, nil
}

func (req *Categories) Search(ctx context.Requester) (int, interface{}) {
	category := ctx.FindParam("category")
	//req.Setup(category, "Blog", false)

	result := []interface{}{}
	pagesize := ctx.FindParam("pagesize")

	_, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Blog.API", "public", category, pagesize)

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, result
}

func (req *Categories) View(ctx context.Requester) (int, interface{}) {
	//req.Setup("article", "Article", false)

	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})

	article := make(map[string]interface{})
	code, err := do.GET(ctx.GetMyToken(), &article, ctx.GetInstanceID(), "Blog.API", "public", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	result["Article"] = article

	comments := []interface{}{}
	code, err = do.GET(ctx.GetMyToken(), &comments, ctx.GetInstanceID(), "Comment.API", "message", "Article", key.String())

	if err != nil && code != 404 {
		log.Println(err)

		return code, err
	}

	result["Comments"] = comments

	return http.StatusOK, result
}
