package resources

func (src *Source) FetchProfileDisplay() (map[string]interface{}, error) {
	obj, err := src.get("cms", "display")

	if err != nil {
		return nil, err
	}
	
	return obj.(map[string]interface{}), nil
}

func (src *Source) FetchCMS(pagesize string) (interface{}, error) {
	return src.get("cms", "content", pagesize)
}
