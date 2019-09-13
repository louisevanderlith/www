package blog

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/husk"
)

type Articles struct {
}

func (c *Articles) Get(ctx context.Requester) (int, interface{}) {
	result := []interface{}{}
	pagesize := "A10"

	_, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Blog.API", "public", pagesize)

	if err != nil {
		log.Println(err)
		return http.StatusBadRequest, err
	}

	return http.StatusOK, result
}

func (c *Articles) Search(ctx context.Requester) (int, interface{}) {
	result := []interface{}{}
	pagesize := ctx.FindParam("pagesize")

	_, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Blog.API", "public", pagesize)

	if err != nil {
		log.Println(err)
		return http.StatusBadRequest, err
	}

	return http.StatusOK, result
}

func (c *Articles) View(ctx context.Requester) (int, interface{}) {
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})

	var article interface{}
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
