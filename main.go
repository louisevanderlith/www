package main

import (
	"flag"
	"github.com/louisevanderlith/kong"
	"github.com/louisevanderlith/www/handles"
	"net/http"
	"time"

	"github.com/louisevanderlith/droxolite"
)

func main() {
	clientId := flag.String("client", "mango.www", "Client ID which will be used to verify this instance")
	clientSecrt := flag.String("secret", "secret", "Client Secret which will be used to authenticate this instance")
	authrty := flag.String("authority", "http://localhost:8086", "Authority Provider's URL")

	flag.Parse()

	scps := []string{
		"comms.messages.create",
		"blog.articles.view",
		"blog.articles.search",
		"comment.messages.view",
		"theme.assets.download",
		"theme.assets.view",
		"artifact.download",
	}

	tkn, err := kong.FetchToken(http.DefaultClient, *authrty, *clientId, *clientSecrt, scps...)

	if err != nil {
		panic(err)
	}

	clms, err := kong.Exchange(http.DefaultClient, tkn, *clientId, *clientSecrt, *authrty+"/info")

	if err != nil {
		panic(err)
	}

	err = droxolite.UpdateTemplate(tkn, clms)

	if err != nil {
		panic(err)
	}

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8091",
		Handler:      handles.SetupRoutes(*clientId, *clientSecrt, *authrty),
	}

	err = srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
