package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/www/controllers"
	"github.com/louisevanderlith/www/controllers/blog"
)

func Setup(e resins.Epoxi, profile string) {
	homeCtrl := &controllers.Home{DefaultProfile: profile}
	e.JoinBundle("/blog", roletype.Unknown, mix.Page, &blog.Articles{}, &blog.Categories{})
	e.JoinPath(e.Router().(*mux.Router), "/{siteName:[a-zA-Z]+}", "index", http.MethodGet, roletype.Unknown, mix.Page, homeCtrl.GetSite)
}
