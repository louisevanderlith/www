package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
)

var (
	config    oauth2.Config
	globToken *oauth2.Token
	Oper service
)

type service struct {
	Profile string
	Host string
	LogoKey string
}

func AssignOperator(profile, host string) {
	Oper = service{
		Profile: profile,
		Host:    fmt.Sprintf(".%s/", host),
		LogoKey: "0`0",
	}
}

func DefineClient(clientId, clientSecret, host, authHost string) {
	config = oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{"blog"},
		RedirectURL:  host + "/oauth2",
		Endpoint: oauth2.Endpoint{
			AuthURL:  authHost + "/authorize",
			TokenURL: authHost + "/token",
		},
	}
}

func AuthCallback(c *gin.Context) {
	c.Request.ParseForm()
	state := c.Request.Form.Get("state")
	if state != "xyz" {
		c.AbortWithError(http.StatusBadRequest, errors.New("state invalid"))
	}

	code := c.Request.Form.Get("code")
	if code == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("code not found"))
	}

	token, err := config.Exchange(context.Background(), code)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	globToken = token

	e := json.NewEncoder(c.Writer)
	e.SetIndent("", "  ")
	e.Encode(token)
}
