package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	stock "github.com/louisevanderlith/stock/api"
	"log"
	"net/http"
)

//GetDefault returns the 'defaultsite'
func Index(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clnt := credConfig.Client(r.Context())
		data, err := stock.FetchClientCategories(clnt, Endpoints["stock"], "A6")

		if err != nil {
			log.Println("Fetch Categories Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		bag := mix.NewDataBag(data)
		bag.SetValue("Title", "Home")
		err = mix.Write(w, fact.Create(r, "Index", "./views/index.html", bag))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
