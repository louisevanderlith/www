package resources

func (src *Source) FetchServices(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("stock", "services", pagesize)

	if err != nil {
		return nil, err
	}

	return res.(map[string]interface{}), nil
}