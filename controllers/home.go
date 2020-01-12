package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPage(c *gin.Context) {
	folioURL := fmt.Sprintf("http://folio:8090/profile/%s", Oper.Profile)
	resp, err := http.Get(folioURL)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	defer resp.Body.Close()

	result := make(map[string]interface{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":      "Home - " + Oper.Profile,
		"Data":       result,
		"Oper":       Oper,
		"HasScript":  true,
		"ScriptName": "index.js",
	})
}
