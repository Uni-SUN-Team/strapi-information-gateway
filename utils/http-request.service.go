package utils

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"time"
	"unisun/api/strapi-information-gateway/constants"
)

func HTTPRequest(url string, method string, payload []byte) *http.Response {
	var request *http.Request
	var err error
	var body *bytes.Buffer
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}
	switch method {
	case constants.GET:
		body = bytes.NewBuffer(nil)
	case constants.POST:
		body = bytes.NewBuffer(payload)
	case constants.PUT:
		body = bytes.NewBuffer(payload)
	case constants.DELETE:
		body = bytes.NewBuffer(nil)
	default:
		body = bytes.NewBuffer(nil)
	}
	if err != nil {
		log.Println("Create request error.", err.Error())
	}
	request, err = http.NewRequest(method, url, body)
	if err != nil {
		log.Println("Client request to "+url+" is not success.", err.Error())
	}
	bearer := "Bearer " + os.Getenv(constants.TOKEN_STRAPI)
	request.Header.Add("Authorization", bearer)
	request.Header.Set("Content-type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		log.Println("Client is error.", err.Error())
	}
	return response
}
