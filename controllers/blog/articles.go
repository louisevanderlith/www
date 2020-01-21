package blog

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/droxo"
	"net/http"

	"github.com/louisevanderlith/husk"
)

func Get(c *gin.Context) {
	pagesize := "A10"

	blogURL := fmt.Sprintf("%sarticles/%s/", droxo.UriBlog, pagesize)
	resp, err := http.Get(blogURL)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	defer resp.Body.Close()

	result := make(map[string]interface{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "articles.html", droxo.Wrap("Articles", result))
}

func Search(c *gin.Context) {
	pagesize := c.Param("pagesize")
	hsh := c.Param("hash")

	blogURL := fmt.Sprintf("%sarticles/%s/%s", droxo.UriBlog, pagesize, hsh)
	resp, err := http.Get(blogURL)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	defer resp.Body.Close()

	result := make(map[string]interface{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	c.HTML(http.StatusOK, "articles.html", droxo.Wrap("Articles", result))
}

func View(c *gin.Context) {
	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	article, err := getArticle(key)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	comments, err := getComments(key)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	result := make(map[string]interface{})
	result["Article"] = article
	result["Comments"] = comments

	c.HTML(http.StatusOK, "articlesview.html", droxo.Wrap("ArticlesView", result))
}

func getArticle(key husk.Key) (map[string]interface{}, error){
	blogURL := fmt.Sprintf("http://blog:8102/article/%s", key)
	resp, err := http.Get(blogURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	article := make(map[string]interface{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&article)

	if err != nil {
		return nil, err
	}

	return article, nil
}

func getComments(key husk.Key) (map[string]interface{}, error) {
	commntURL := fmt.Sprintf("%smessage/Article/%s", droxo.UriComment, key)
	resp, err := http.Get(commntURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	comments := make(map[string]interface{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&comments)

	if err != nil {
		return nil, err
	}

	return comments, nil
}