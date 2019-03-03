package main

import (
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"
	"github.com/louisevanderlith/www/routers"
)

func main() {
	// Register with router
	mode := os.Getenv("RUNMODE")
	log.Println("Running in ", mode)
	name := beego.BConfig.AppName
	srv := mango.NewService(mode, name, enums.APP)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(srv)
		beego.Run()
	}
}
