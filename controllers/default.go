package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/xontrols"
)

type DefaultController struct {
	xontrols.UICtrl
}

//GetDefault returns the 'defaultsite'
func (c *DefaultController) GetDefault() {
	result := make(map[string]interface{})
	code, err := droxolite.DoGET("", &result, c.Settings.InstanceID, "Folio.API", "profile", c.Settings.Name)

	if err != nil {
		c.Serve(code, err, nil)
		return
	}

	c.Setup("default", "Home", true)
	c.CreateTopMenu(getHomeMenu())

	c.Serve(http.StatusOK, nil, result)
}

func (c *DefaultController) GetSite() {
	siteName := c.FindParam("siteName")

	result := make(map[string]interface{})
	code, err := droxolite.DoGET("", &result, c.Settings.InstanceID, "Folio.API", "profile", siteName)

	if err != nil {
		c.Serve(code, err, nil)
		return
	}

	pageTitle := "Home"
	dataObj, ok := result["Data"].(map[string]interface{})

	if ok {
		pageTitle = dataObj["Title"].(string)
	}

	c.Setup("default", pageTitle, true)
	c.CreateTopMenu(getHomeMenu())
	c.Serve(http.StatusOK, nil, result)
}

func getHomeMenu() *bodies.Menu {
	result := bodies.NewMenu("/home")

	result.AddItem("#portfolio", "Portfolio", "home fa fa-star", nil)
	result.AddItem("#aboutus", "About Us", "home fa fa-info", nil)
	result.AddItem("#contact", "Contact", "home fa fa-phone", nil)
	result.AddItem("/blogs/A10", "Blog", "home fa fa-phone", nil)

	return result
}
