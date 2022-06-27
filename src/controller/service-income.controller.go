package controller

import (
	"log"
	"net/http"

	"unisun/api/strapi-information-gateway/src/model"
	"unisun/api/strapi-information-gateway/src/ports"

	"github.com/gin-gonic/gin"
)

type ControllerServiceIncomeAdapter struct {
	Service ports.ServiceIncome
}

func NewControllerServiceIncomeAdapter(service ports.ServiceIncome) *ControllerServiceIncomeAdapter {
	return &ControllerServiceIncomeAdapter{
		Service: service,
	}
}

// StrapiInformationGatewayHandler godoc
// @summary      Strapi Information Gateway
// @description  Strapi Information Gateway for the service
// @id           StrapiInformationGatewayHandler
// @produce      json
// @success     200  {object}  model.Response  "OK"
// @failure		400  {object}  model.Response  "FAIL"
// @router       /strapi [post]
func (srv *ControllerServiceIncomeAdapter) StrapiGateway(c *gin.Context) {
	var body = model.Payload{}
	var responseStruct = model.Response{}
	if err := c.ShouldBindJSON(&body); err != nil {
		responseStruct.Error = err.Error()
		responseStruct.Status = false
		responseStruct.Payload = ""
		c.AbortWithStatusJSON(http.StatusBadRequest, responseStruct)
		log.Panic("Bind json body is error.", err)
		return
	}
	response, err := srv.Service.CallStrapi(body)
	if err != nil {
		responseStruct.Error = err.Error()
		responseStruct.Status = false
		responseStruct.Payload = string(response)
		c.AbortWithStatusJSON(http.StatusBadRequest, responseStruct)
		log.Panic("Call strapi is error!.", err)
		return
	}
	responseStruct.Error = ""
	responseStruct.Status = true
	responseStruct.Payload = string(response)
	c.AbortWithStatusJSON(http.StatusOK, responseStruct)
}
