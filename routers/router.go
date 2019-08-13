package routers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/droxolite/routing"
	"github.com/louisevanderlith/www/controllers"
)

func Setup(e *droxolite.Epoxy) {
	//Home
	homeCtrl := &controllers.Home{}
	homeGroup := routing.NewInterfaceBundle("", roletype.Unknown, homeCtrl)
	homeGroup.AddRoute("Profile", "/{siteName:[a-zA-Z]+}", http.MethodGet, roletype.Unknown, homeCtrl.GetSite)
	e.AddNamedGroup("Home", homeGroup)

	//Blog
	blogCtrl := &controllers.Blog{}
	blogGroup := routing.NewInterfaceBundle("Blog", roletype.Unknown, blogCtrl)
	blogGroup.AddRoute("Articles by Category", "/articles/{category:[a-zA-Z]+}/{pagesize:[A-Z][0-9]+}", http.MethodGet, roletype.Unknown, blogCtrl.SearchByCategory)

	e.AddNamedGroup("Blog", blogGroup)
}
