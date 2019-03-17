package controllers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type DefaultController struct {
	control.UIController
}

func NewDefaultCtrl(ctrlMap *control.ControllerMap) *DefaultController {
	result := &DefaultController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *DefaultController) GetDefault() {
	c.Setup("default")
	c.CreateTopMenu(getTopMenu())
	siteName := beego.AppConfig.String("defaultsite")

	result := make(map[string]interface{})
	fail, err := mango.DoGET(&result, c.GetInstanceID(), "Folio.API", "profile", siteName)

	if err != nil {
		c.Serve(nil, err)
		return
	}

	if fail != nil {
		c.Serve(nil, fail)
		return
	}

	c.Serve(result, nil)
}

func (c *DefaultController) GetSite() {
	c.Setup("default")
	c.CreateTopMenu(getTopMenu())
	siteName := c.Ctx.Input.Param(":siteName")

	result := make(map[string]interface{})
	fail, err := mango.DoGET(&result, c.GetInstanceID(), "Folio.API", "profile", siteName)

	if err != nil {
		c.Serve(nil, err)
		return
	}

	if fail != nil {
		c.Serve(nil, fail)
		return
	}

	c.Serve(result, nil)
}

func getTopMenu() *control.Menu {
	result := control.NewMenu("/home")

	result.AddItem("#portfolio", "Portfolio", "home fa fa-star", nil)
	result.AddItem("#aboutus", "About Us", "home fa fa-info", nil)
	result.AddItem("#contact", "Contact", "home fa fa-phone", nil)

	return result
}
