package controllers

import (
	"log"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"

	"github.com/louisevanderlith/mango/control"
)

type BlogController struct {
	control.UIController
}

func NewBlogCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) *BlogController {
	result := &BlogController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *BlogController) Get() {
	c.Setup("blog", "Blog", false)
	c.CreateSideMenu(getBlogMenu())

	result := []interface{}{}
	pagesize := c.Ctx.Input.Param(":pagesize")

	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Blog.API", "article", "all", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(nil, err)
		return
	}

	c.Serve(result, nil)
}

func (c *BlogController) GetByCategory() {
	c.Setup("blog", "Blog", false)
	c.CreateSideMenu(getBlogMenu())

	result := []interface{}{}
	category := c.Ctx.Input.Param(":category")
	pagesize := c.Ctx.Input.Param(":pagesize")

	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Blog.API", "article", "all", category, pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(nil, err)
		return
	}

	c.Serve(result, nil)
}

func (c *BlogController) GetArticle() {
	c.Setup("article", "Article", false)
	c.CreateSideMenu(getBlogMenu())
	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
		return
	}

	result := make(map[string]interface{})

	article := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &article, c.GetInstanceID(), "Blog.API", "article", key.String())

	if err != nil {
		log.Println(err)
		c.Serve(nil, err)
		return
	}

	result["Article"] = article

	comments := []interface{}{}
	code, err := mango.DoGET(c.GetMyToken(), &comments, c.GetInstanceID(), "Comment.API", "message", "Article", key.String())

	if err != nil && code != 404 {
		c.Serve(nil, err)
		return
	}

	result["Comments"] = comments

	c.Serve(result, nil)
}

func getBlogMenu() *control.Menu {
	result := control.NewMenu("/home")

	result.AddItem("#", "Categories", "fa fa-cirlce", categoryChlidren("/categorie"))

	return result
}

func categoryChlidren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("/blogs/motoring/A10", "Motoring", "fa fa-car", nil)
	children.AddItem("/blogs/technology/A10", "Technology", "fa fa-microchip", nil)

	return children
}
