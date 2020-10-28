package main

import (
	"flag"
	"github.com/louisevanderlith/www/handles"
	"net/http"
	"time"
)

func main() {
	host := flag.String("host", "http://127.0.0.1:8091", "This application's URL")
	clientId := flag.String("client", "mango.www", "Client ID which will be used to verify this instance")
	clientSecret := flag.String("secret", "secret", "Client Secret which will be used to authenticate this instance")
	issuer := flag.String("issuer", "http://127.0.0.1:8080/auth/realms/mango", "OIDC Provider's URL")
	theme := flag.String("theme", "http://127.0.0.1:8093", "Theme URL")
	folio := flag.String("folio", "http://127.0.0.1:8090", "Folio URL")
	blog := flag.String("blog", "http://127.0.0.1:8102", "Blog URL")
	stock := flag.String("stock", "http://127.0.0.1:8101", "Stock URL")
	comms := flag.String("comms", "http://127.0.0.1:8085", "Communications URL")
	flag.Parse()

	ends := map[string]string{
		"issuer": *issuer,
		"theme":  *theme,
		"folio":  *folio,
		"blog":   *blog,
		"stock":  *stock,
		"comms":  *comms,
	}

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8091",
		Handler:      handles.SetupRoutes(*host, *clientId, *clientSecret, ends),
	}

	err := srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
