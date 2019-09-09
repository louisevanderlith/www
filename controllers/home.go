package controllers

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/do"
)

type Home struct {
	DefaultProfile string
}

//GetDefault returns the 'defaultsite'
func (c *Home) Get(ctx context.Requester) (int, interface{}) {
	result := make(map[string]interface{})
	log.Println(c.DefaultProfile)
	code, err := do.GET("", &result, ctx.GetInstanceID(), "Folio.API", "profile", c.DefaultProfile)

	if err != nil {
		return code, err
	}

	log.Println(result)

	return http.StatusOK, result
}

func (c *Home) GetSite(ctx context.Requester) (int, interface{}) {
	siteName := ctx.FindParam("siteName")

	result := make(map[string]interface{})
	code, err := do.GET("", &result, ctx.GetInstanceID(), "Folio.API", "profile", siteName)

	if err != nil {
		return code, err
	}

	return http.StatusOK, result
}
