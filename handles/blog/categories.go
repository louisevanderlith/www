package blog

/*
func (req *Categories) Get(ctx context.Requester) (int, interface{}) {
	//Show categories...
	//req.Setup("categories", "Blog Categories", false)
	var categories []string
	categories = append(categories, "Motoring", "Technology")

	return http.StatusOK, categories
}

func (req *Categories) SearchCategory(ctx context.Requester) (int, interface{}) {
	category := drx.FindParam(r,"category")

	result := []interface{}{}
	pagesize := drx.FindParam(r,"pagesize")

	_, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Blog.API", "public", category, pagesize)

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, result
}

func (req *Categories) Search(ctx context.Requester) (int, interface{}) {
	category := drx.FindParam(r,"category")

	result := []interface{}{}
	pagesize := drx.FindParam(r,"pagesize")

	_, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Blog.API", "public", category, pagesize)

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, result
}

func (req *Categories) View(ctx context.Requester) (int, interface{}) {
	key, err := keys.ParseKey(drx.FindParam(r,"key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})

	var article interface{}
	code, err := do.GET(ctx.GetMyToken(), &article, ctx.GetInstanceID(), "Blog.API", "public", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	result["Article"] = article

	comments := []interface{}{}
	code, err = do.GET(ctx.GetMyToken(), &comments, ctx.GetInstanceID(), "Comment.API", "message", "Article", key.String())

	if err != nil && code != 404 {
		log.Println(err)

		return code, err
	}

	result["Comments"] = comments

	return http.StatusOK, result
}
*/
