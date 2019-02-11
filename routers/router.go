package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/www/controllers"
)

func Setup(s *mango.Service) {
	ctrlmap := control.CreateControlMap(s)
	beego.Router("/", controllers.NewDefaultCtrl(ctrlmap))
	beego.Router("/:siteName", controllers.NewDefaultCtrl(ctrlmap))
}
