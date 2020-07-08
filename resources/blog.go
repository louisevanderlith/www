package resources

func (src *Source) FetchArticles(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("blog", "articles", pagesize)

	if err != nil {
		return nil, err
	}

	return res.(map[string]interface{}), nil
}

func (src *Source) FetchArticle(key string) (interface{}, error) {
	return src.get("blog", "articles", key)
}
