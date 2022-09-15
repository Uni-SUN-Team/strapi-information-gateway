package services

import (
	"fmt"
	"strings"
	"unisun/api/unisun-strapi-inquiry/src/configs/environment"
	"unisun/api/unisun-strapi-inquiry/src/models"

	"github.com/narawichsaphimarn/handlercontrol/interfaces"
)

type httpClient struct {
	HttpClient interfaces.HttpRequest
}

func New(httpRequest interfaces.HttpRequest) *httpClient {
	return &httpClient{
		HttpClient: httpRequest,
	}
}

func (srv *httpClient) CallStrapi(payload models.Payload) ([]byte, error) {
	endpointEnv := environment.ENV.Endpoint
	url := strings.Join([]string{endpointEnv.Strapi.Host, payload.Path}, "")
	response, err := srv.HttpClient.ClientRequest(url, payload.Method, payload.Body)
	if err != nil {
		return nil, fmt.Errorf("%s{%v}", "Error step call strapi for get value.Error msg : ", err)
	}
	return response, nil
}
