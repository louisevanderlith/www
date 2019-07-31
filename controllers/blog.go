package controllers

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
)

type BlogController struct {
	xontrols.UICtrl
}

func (c *BlogController) Get() {
	c.Setup("blog", "Blog", false)
	c.CreateSideMenu(getBlogMenu())

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

func (c *BlogController) GetByCategory() {
	c.Setup("blog", "Blog", false)
	c.CreateSideMenu(getBlogMenu())

	result := []interface{}{}
	category := c.FindParam("category")
	pagesize := c.FindParam("pagesize")

	_, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Blog.API", "article", "all", category, pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(http.StatusInternalServerError, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *BlogController) GetArticle() {
	c.Setup("article", "Article", false)
	c.CreateSideMenu(getBlogMenu())
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

func getBlogMenu() *bodies.Menu {
	result := bodies.NewMenu()

	result.AddItem("#", "Categories", "fa fa-cirlce", categoryChlidren())

	return result
}

func categoryChlidren() *bodies.Menu {
	children := bodies.NewMenu()
	children.AddItem("/blogs/motoring/A10", "Motoring", "fa fa-car", nil)
	children.AddItem("/blogs/technology/A10", "Technology", "fa fa-microchip", nil)

	return children
}
