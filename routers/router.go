package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/www/controllers"
)

func Setup(s *mango.Service) {
	ctrlmap := control.CreateControlMap(s)

	siteName := beego.AppConfig.String("defaultsite")
	theme, err := mango.GetDefaultTheme(ctrlmap.GetInstanceID(), siteName)

	if err != nil {
		panic(err)
	}

	deftCtrl := controllers.NewDefaultCtrl(ctrlmap, theme)

	beego.Router("/", deftCtrl, "get:GetDefault")
	beego.Router("/:siteName", deftCtrl, "get:GetSite")

	blogCtrl := controllers.NewBlogCtrl(ctrlmap, theme)

	beego.Router("/blog/:pagesize", blogCtrl, "get:Get")
	beego.Router("/article/:key", blogCtrl, "get:GetArticle")
}
