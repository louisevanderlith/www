package main

import (
	"os"
	"path"
	"strconv"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/droxolite/element"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/servicetype"
	"github.com/louisevanderlith/www/controllers"
	"github.com/louisevanderlith/www/routers"
)

func main() {
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	host := os.Getenv("HOST")
	httpport, _ := strconv.Atoi(os.Getenv("HTTPPORT"))
	profile := os.Getenv("PROFILE")
	appName := os.Getenv("APPNAME")
	pubPath := path.Join(keyPath, pubName)

	// Register with router
	srv := bodies.NewService(appName, profile, pubPath, httpport, servicetype.APP)

	routr, err := do.GetServiceURL("", "Router.API", false)

	if err != nil {
		panic(err)
	}

	err = srv.Register(routr)

	if err != nil {
		panic(err)
	}

	err = droxolite.UpdateTheme(srv.ID)

	if err != nil {
		panic(err)
	}

	theme, err := element.GetDefaultTheme(host, srv.ID, profile)

	if err != nil {
		panic(err)
	}

	err = theme.LoadTemplate("./views", "master.html")

	if err != nil {
		panic(err)
	}
	homeCtrl := &controllers.Home{DefaultProfile: profile}
	poxy := resins.NewColourEpoxy(srv, theme, "", homeCtrl.Get)
	routers.Setup(poxy, profile)

	err = droxolite.Boot(poxy)

	if err != nil {
		panic(err)
	}
}
