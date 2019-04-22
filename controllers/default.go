package controllers

import (
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type DefaultController struct {
	control.UIController
	SiteName string
}

func NewDefaultCtrl(ctrlMap *control.ControllerMap, theme mango.ThemeSetting) *DefaultController {
	result := &DefaultController{
		SiteName: theme.Name,
	}

	result.SetTheme(theme)
	result.SetInstanceMap(ctrlMap)

	return result
}

//GetDefault returns the 'defaultsite'
func (c *DefaultController) GetDefault() {
	result := make(map[string]interface{})
	err := mango.DoGET(&result, c.GetInstanceID(), "Folio.API", "profile", c.SiteName)

	if err != nil {
		c.Serve(nil, err)
		return
	}

	c.Setup("default", "Home", true)
	c.CreateTopMenu(getTopMenu())
	c.Serve(result, nil)
}

func (c *DefaultController) GetSite() {
	siteName := c.Ctx.Input.Param(":siteName")

	result := make(map[string]interface{})
	err := mango.DoGET(&result, c.GetInstanceID(), "Folio.API", "profile", siteName)

	if err != nil {
		c.Serve(nil, err)
		return
	}

	pageTitle := "Home"
	dataObj, ok := result["Data"].(map[string]interface{})

	if ok {
		pageTitle = dataObj["Title"].(string)
	}

	c.Setup("default", pageTitle, true)
	c.CreateTopMenu(getTopMenu())
	c.Serve(result, nil)
}

func getTopMenu() *control.Menu {
	result := control.NewMenu("/home")

	result.AddItem("#portfolio", "Portfolio", "home fa fa-star", nil)
	result.AddItem("#aboutus", "About Us", "home fa fa-info", nil)
	result.AddItem("#contact", "Contact", "home fa fa-phone", nil)

	return result
}
