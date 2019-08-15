package blog

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
)

type Categories struct {
	xontrols.UICtrl
}

func (req *Categories) Default() {
	//Show categories...
	req.Setup("categories", "Blog Categories", false)
	req.Serve(http.StatusOK, nil, nil)
}

func (req *Categories) Search() {
	category := req.FindParam("category")
	req.Setup(category, "Blog", false)

	result := []interface{}{}
	pagesize := req.FindParam("pagesize")

	_, err := droxolite.DoGET(req.GetMyToken(), &result, req.Settings.InstanceID, "Blog.API", "article", "all", category, pagesize)

	if err != nil {
		log.Println(err)
		req.Serve(http.StatusInternalServerError, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, result)
}

func (req *Categories) View() {
	req.Setup("article", "Article", false)

	key, err := husk.ParseKey(req.FindParam("key"))

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	result := make(map[string]interface{})

	article := make(map[string]interface{})
	code, err := droxolite.DoGET(req.GetMyToken(), &article, req.Settings.InstanceID, "Blog.API", "article", key.String())

	if err != nil {
		log.Println(err)
		req.Serve(code, err, nil)
		return
	}

	result["Article"] = article

	comments := []interface{}{}
	code, err = droxolite.DoGET(req.GetMyToken(), &comments, req.Settings.InstanceID, "Comment.API", "message", "Article", key.String())

	if err != nil && code != 404 {
		log.Println(err)
		req.Serve(code, err, nil)
		return
	}

	result["Comments"] = comments

	req.Serve(http.StatusOK, nil, result)
}
