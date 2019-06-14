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
	//category

	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Blog.API", "article", "all", pagesize)

	c.Serve(result, err)
}

func (c *BlogController) GetArticle() {
	c.Setup("article", "Article", false)
	c.CreateSideMenu(getBlogMenu())
	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Blog.API", "article", key.String())

	if err != nil {
		log.Println(err)
	}

	c.Serve(result, err)
}

func getBlogMenu() *control.Menu {
	result := control.NewMenu("/home")

	result.AddItem("", "/", "Home", "fa fa-house", nil)
	result.AddItem("", "#", "Categories", "fa fa-cirlce", categoryChlidren("/categorie"))

	return result
}

func categoryChlidren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("","/blog/cars", "Cars Blog", "fa fa-car", nil)
	children.AddItem("","/blog/tech", "Technology Blog", "fa fa-news", nil)

	return children
}
