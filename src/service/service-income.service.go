package service

import (
	"io/ioutil"
	"log"
	"os"
	"unisun/api/strapi-information-gateway/src/model"
	"unisun/api/strapi-information-gateway/src/utils"
)

func CallStrapi(payload model.Payload) ([]byte, error) {
	var url = os.Getenv("HOST_STRAPI") + payload.Path
	response := utils.HTTPRequest(url, payload.Method, payload.Body)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Read response from request", err.Error())
	} else {
		err = nil
	}
	defer response.Body.Close()
	return body, err
}
