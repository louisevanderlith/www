package logic

import (
	"encoding/json"
	"log"

	folio "github.com/louisevanderlith/folio/core"
	"github.com/louisevanderlith/mango"
)

var uploadURL string

func GetProfileSite(instanceID, name string) (folio.Profile, error) {
	result := folio.Profile{}
	if name == "" {
		name = "avosa"
	}

	resp, err := mango.GETMessage(instanceID, "Folio.API", "profile", name)

	if err != nil {
		return result, err
	}

	if resp.Failed() {
		return result, resp
	}

	switch resp.Data.(type) {
	case map[string]interface{}:
		dirty, err := json.Marshal(resp.Data)

		if err != nil {
			return result, err
		}

		err = json.Unmarshal(dirty, &result)

		if err != nil {
			return result, err
		}
	case folio.Profile:
		result = resp.Data.(folio.Profile)
	default:
		log.Printf("Dont Know: %v\n", resp)
	}

	//result = resp.Data.(folio.Profile)
	//result.setImageURLs(instanceID)

	return result, nil
}

/*
func (obj *BasicSite) setImageURLs(instanceID string) {
	if uploadURL == "" {
		setUploadURL(instanceID)
	}

	obj.ImageURL = uploadURL + strconv.FormatInt(obj.ImageID, 10)

	for i := 0; i < len(obj.PortfolioItems); i++ {
		row := &obj.PortfolioItems[i]
		row.ImageURL = uploadURL + strconv.FormatInt(row.ImageID, 10)
	}

	for i := 0; i < len(obj.Headers); i++ {
		row := &obj.Headers[i]
		row.ImageURL = uploadURL + strconv.FormatInt(row.ImageID, 10)
	}
}

func setUploadURL(instanceID string) {
	url, err := mango.GetServiceURL(instanceID, "Artifact.API", true)

	if err != nil {
		log.Print("setImageURLs:", err)
	}

	uploadURL = url + "v1/upload/file/"
}
*/
