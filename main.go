package main

import (
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/droxo"
	"github.com/louisevanderlith/www/controllers"
	"github.com/louisevanderlith/www/controllers/blog"
	"os"
)

func main() {
	prof := os.Getenv("PROFILE")

	if len(prof) == 0 {
		panic("invalid profile")
	}

	host := os.Getenv("HOST")

	droxo.AssignOperator(prof, host)
	droxo.DefineClient("www", "wwwsecret", host, "http://oauth2." + host)
	//Download latest Theme
	err := droxo.UpdateTheme("http://theme:8093")

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	tmpl, err := droxo.LoadTemplates("./views")

	if err != nil {
		panic(err)
	}

	r.HTMLRender = tmpl

	r.GET("/", controllers.IndexPage)
	r.GET("/blog", blog.Get)
	r.GET("/blog/:pagesize/*hash", blog.Search)
	r.GET("/article/:key", blog.View)
	//r.POST("/oauth2", droxo.AuthCallback)

	err = r.Run(":8091")

	if err != nil {
		panic(err)
	}
}
