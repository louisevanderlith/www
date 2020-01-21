package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/louisevanderlith/droxo"

	"github.com/gin-gonic/gin"
)

func IndexPage(c *gin.Context) {
	folioURL := fmt.Sprintf("%sprofile/%s", droxo.UriFolio, droxo.Oper.Profile)
	resp, err := http.Get(folioURL)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	defer resp.Body.Close()

	result := make(map[string]interface{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.HTML(http.StatusOK, "index.html", droxo.Wrap("Index", result))
}
