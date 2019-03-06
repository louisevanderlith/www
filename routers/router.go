package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/www/controllers"
)

func Setup(s *mango.Service) {
	ctrlmap := control.CreateControlMap(s)
	deftCtrl := controllers.NewDefaultCtrl(ctrlmap)

	beego.Router("/", deftCtrl, "get:GetDefault")
	beego.Router("/:siteName", deftCtrl, "get:GetSite")
}
