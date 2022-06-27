package service

import (
	"io/ioutil"
	"unisun/api/strapi-information-gateway/src/model"
	"unisun/api/strapi-information-gateway/src/ports"

	"github.com/spf13/viper"
)

type UtilsService struct {
	UtilsService ports.HTTPService
}

func New(utilsService ports.HTTPService) *UtilsService {
	return &UtilsService{
		UtilsService: utilsService,
	}
}

func (svr *UtilsService) CallStrapi(payload model.Payload) ([]byte, error) {
	var url = viper.GetString("endpoint.strapi.host") + payload.Path
	response, err := svr.UtilsService.HTTPRequest(url, payload.Method, payload.Body)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return body, nil
}
