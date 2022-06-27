package utils

import (
	"bytes"
	"net/http"
	"time"
	"unisun/api/strapi-information-gateway/src/constants"
)

type Service struct {
	AccessToken string
}

func New(accessToken string) *Service {
	return &Service{
		AccessToken: accessToken,
	}
}

func (svr *Service) HTTPRequest(url string, method string, payload []byte) (*http.Response, error) {
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
		return nil, err
	}
	request, err = http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	bearer := "Bearer " + svr.AccessToken
	request.Header.Add("Authorization", bearer)
	request.Header.Set("Content-type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
