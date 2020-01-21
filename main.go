package main

import (
	"encoding/json"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/www/controllers"
	"github.com/louisevanderlith/www/controllers/blog"
	"github.com/louisevanderlith/droxo"
	"io"
	"net/http"
	"os"
	"path/filepath"
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
	err := updateTheme("http://theme:8093")

	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.HTMLRender = loadTemplates("./views")

	r.GET("/", controllers.IndexPage)
	r.GET("/blog", blog.Get)
	r.GET("/blog/:pagesize/*hash", blog.Search)
	r.GET("/article/:key", blog.View)
	r.POST("/oauth2", droxo.AuthCallback)

	err = r.Run(":8091")

	if err != nil {
		panic(err)
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/_shared/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/*.html")

	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		base := filepath.Base(include)

		r.AddFromFiles(base, files...)
	}
	return r
}

func updateTheme(url string) error {
	resp, err := http.Get(url + "/asset/html")

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var items []string
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&items)

	if err != nil {
		return err
	}

	for _, v := range items {
		err = downloadFile(url, v)

		if err != nil {
			return err
		}
	}

	return nil
}

func downloadFile(url, templ string) error {
	resp, err := http.Get(url + "/asset/html/" + templ)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create("/views/_shared/" + templ)

	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}