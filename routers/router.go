package routers

import (
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/www/controllers"
)

func Setup(e *droxolite.Epoxy) {
	//Home
	deftCtrl := &controllers.DefaultController{}
	deftGroup := droxolite.NewRouteGroup("", deftCtrl)
	deftGroup.AddRoute("/", "GET", roletype.Unknown, deftCtrl.GetDefault)
	deftGroup.AddRoute("/{siteName:[a-zA-Z]+}", "GET", roletype.Unknown, deftCtrl.GetSite)
	e.AddGroup(deftGroup)

	//Blog
	blogCtrl := &controllers.BlogController{}
	blogGroup := droxolite.NewRouteGroup("blogs", blogCtrl)
	blogGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Unknown, blogCtrl.Get)
	blogGroup.AddRoute("/{category:[a-zA-Z]+}/{pagesize:[A-Z][0-9]+}", "GET", roletype.Unknown, blogCtrl.GetByCategory)
	blogGroup.AddRoute("/article/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, blogCtrl.GetArticle)
	e.AddGroup(blogGroup)
	/*ctrlmap := control.CreateControlMap(s)

	siteName := beego.AppConfig.String("defaultsite")
	theme, err := mango.GetDefaultTheme(ctrlmap.GetInstanceID(), siteName)

	if err != nil {
		panic(err)
	}

	deftCtrl := controllers.NewDefaultCtrl(ctrlmap, theme)

	beego.Router("/", deftCtrl, "get:GetDefault")
	beego.Router("/:siteName", deftCtrl, "get:GetSite")

	blogCtrl := controllers.NewBlogCtrl(ctrlmap, theme)

	beego.Router("/blogs/:pagesize", blogCtrl, "get:Get")
	beego.Router("/blogs/:category/:pagesize", blogCtrl, "get:GetByCategory")
	beego.Router("/article/:key", blogCtrl, "get:GetArticle")*/
}
