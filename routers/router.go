package routers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/droxolite/routing"
	"github.com/louisevanderlith/www/controllers"
	"github.com/louisevanderlith/www/controllers/blog"
)

func Setup(e resins.Epoxi, profile string) {
	//Home
	homeCtrl := &controllers.Home{DefaultProfile: profile}
	homeGroup := routing.NewInterfaceBundle("", roletype.Unknown, homeCtrl)
	homeGroup.AddRoute("Profile", "/{siteName:[a-zA-Z]+}", http.MethodGet, roletype.Unknown, homeCtrl.GetSite)
	e.AddGroup(homeGroup)

	//Blog
	//blogGroup.AddRoute("Articles by Category", "/articles/{category:[a-zA-Z]+}/{pagesize:[A-Z][0-9]+}", http.MethodGet, roletype.Unknown, blogCtrl.SearchByCategory)
	blogGroup := routing.NewInterfaceBundle("Blog", roletype.Unknown, &blog.Articles{}, &blog.Categories{})
	e.AddGroup(blogGroup)
}
