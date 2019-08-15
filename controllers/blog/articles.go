package blog

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
)

type Articles struct {
	xontrols.UICtrl
}

func (c *Articles) Default() {
	c.Setup("blog", "Blog", false)

	result := []interface{}{}
	pagesize := "A10"

	_, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Blog.API", "article", "all", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(http.StatusBadRequest, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *Articles) Search() {
	c.Setup("blog", "Blog", false)

	result := []interface{}{}
	pagesize := c.FindParam("pagesize")

	_, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Blog.API", "article", "all", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(http.StatusBadRequest, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *Articles) View() {
	c.Setup("article", "Article", false)

	key, err := husk.ParseKey(c.FindParam("key"))

	if err != nil {
		c.Serve(http.StatusBadRequest, err, nil)
		return
	}

	result := make(map[string]interface{})

	article := make(map[string]interface{})
	code, err := droxolite.DoGET(c.GetMyToken(), &article, c.Settings.InstanceID, "Blog.API", "article", key.String())

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	result["Article"] = article

	comments := []interface{}{}
	code, err = droxolite.DoGET(c.GetMyToken(), &comments, c.Settings.InstanceID, "Comment.API", "message", "Article", key.String())

	if err != nil && code != 404 {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	result["Comments"] = comments

	c.Serve(http.StatusOK, nil, result)
}
