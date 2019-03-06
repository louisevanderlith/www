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

	resp, err := mango.GETMessage(c.GetInstanceID(), "Folio.API", "profile", siteName)

	if err != nil {
		c.Serve(nil, err)
		return
	}

	if resp.Failed() {
		c.Serve(nil, resp)
		return
	}

	c.Serve(resp.Data, err)
}

func (c *DefaultController) GetSite() {
	c.Setup("default")
	c.CreateTopMenu(getTopMenu())
	siteName := c.Ctx.Input.Param(":siteName")

	resp, err := mango.GETMessage(c.GetInstanceID(), "Folio.API", "profile", siteName)

	if err != nil {
		c.Serve(nil, err)
		return
	}

	if resp.Failed() {
		c.Serve(nil, resp)
		return
	}

	c.Serve(resp.Data, err)
}

func getTopMenu() *control.Menu {
	result := control.NewMenu("/home")

	result.AddItem("#portfolio", "Portfolio", "home gome fa-home", nil)
	result.AddItem("#aboutus", "About Us", "home gome fa-home", nil)
	result.AddItem("#contact", "Contact", "home gome fa-home", nil)

	return result
}
