package blog

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

type Articles struct {
}

func (c *Articles) Default(ctx context.Contexer) (int, interface{}) {
	//c.Setup("blog", "Blog", false)

	result := []interface{}{}
	pagesize := "A10"

	_, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Blog.API", "article", "all", pagesize)

	if err != nil {
		log.Println(err)
		return http.StatusBadRequest, err
	}

	return http.StatusOK, result
}

func (c *Articles) Search(ctx context.Contexer) (int, interface{}) {
	//c.Setup("blog", "Blog", false)

	result := []interface{}{}
	pagesize := ctx.FindParam("pagesize")

	_, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Blog.API", "article", "all", pagesize)

	if err != nil {
		log.Println(err)
		return http.StatusBadRequest, err
	}

	return http.StatusOK, result
}

func (c *Articles) View(ctx context.Contexer) (int, interface{}) {
	//c.Setup("article", "Article", false)

	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})

	article := make(map[string]interface{})
	code, err := droxolite.DoGET(ctx.GetMyToken(), &article, ctx.GetInstanceID(), "Blog.API", "article", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	result["Article"] = article

	comments := []interface{}{}
	code, err = droxolite.DoGET(ctx.GetMyToken(), &comments, ctx.GetInstanceID(), "Comment.API", "message", "Article", key.String())

	if err != nil && code != 404 {
		log.Println(err)
		return code, err
	}

	result["Comments"] = comments

	return http.StatusOK, result
}
