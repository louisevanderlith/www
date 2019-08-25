package controllers

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/context"
)

type Home struct {
	DefaultProfile string
}

//GetDefault returns the 'defaultsite'
func (c *Home) Default(ctx context.Contexer) (int, interface{}) {
	result := make(map[string]interface{})
	log.Println(c.DefaultProfile)
	code, err := droxolite.DoGET("", &result, ctx.GetInstanceID(), "Folio.API", "profile", c.DefaultProfile)

	if err != nil {
		return code, err
	}

	log.Println(result)
	//c.Setup("default", "Home", true)
	//c.CreateTopMenu(getHomeMenu())

	return http.StatusOK, result
}

func (c *Home) GetSite(ctx context.Contexer) (int, interface{}) {
	siteName := ctx.FindParam("siteName")

	result := make(map[string]interface{})
	code, err := droxolite.DoGET("", &result, ctx.GetInstanceID(), "Folio.API", "profile", siteName)

	if err != nil {
		return code, err
	}

	//pageTitle := "Home"
	//dataObj, ok := result["Data"].(map[string]interface{})

	/*if ok {
		pageTitle = dataObj["Title"].(string)
	}*/

	//c.Setup("default", pageTitle, true)
	//c.CreateTopMenu(getHomeMenu())
	return http.StatusOK, result
}

func getHomeMenu() bodies.Menu {
	result := bodies.NewMenu()

	//result.AddItem("#portfolio", "Portfolio", "home fa fa-star", nil)
	//result.AddItem("#aboutus", "About Us", "home fa fa-info", nil)
	//result.AddItem("#contact", "Contact", "home fa fa-phone", nil)
	//result.AddItem("/blogs/A10", "Blog", "home fa fa-phone", nil)

	return *result
}
