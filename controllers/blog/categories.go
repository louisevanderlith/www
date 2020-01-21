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

func (req *Categories) Get(c *gin.Context) {
	//Show categories...
	//req.Setup("categories", "Blog Categories", false)
	var categories []string
	categories = append(categories, "Motoring", "Technology")

	return http.StatusOK, categories
}

func (req *Categories) SearchCategory(ctx context.Requester) (int, interface{}) {
	category := c.Param("category")

	result := []interface{}{}
	pagesize := c.Param("pagesize")

	_, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Blog.API", "public", category, pagesize)

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, result
}

func (req *Categories) Search(c *gin.Context) {
	category := c.Param("category")

	result := []interface{}{}
	pagesize := c.Param("pagesize")

	_, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Blog.API", "public", category, pagesize)

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, result
}

func (req *Categories) View(c *gin.Context) {
	key, err := husk.ParseKey(c.Param("key"))

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
