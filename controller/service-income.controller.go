package controller

import (
	"log"
	"net/http"
	"unisun/api/strapi-information-gateway/model"
	"unisun/api/strapi-information-gateway/service"

	"github.com/gin-gonic/gin"
)

// StrapiInformationGatewayHandler godoc
// @summary      Strapi Information Gateway
// @description  Strapi Information Gateway for the service
// @id           StrapiInformationGatewayHandler
// @produce      json
// @response     200  {object}  model.Response  "OK"
// @router       /strapi [get]
func StrapiGateway(c *gin.Context) {
	var body = model.Payload{}
	var responseStruct = model.Response{}
	if err := c.ShouldBindJSON(&body); err != nil {
		responseStruct.Error = err.Error()
		responseStruct.Status = false
		responseStruct.Payload = ""
		c.JSON(http.StatusBadRequest, responseStruct)
		return
	}
	response, error := service.CallStrapi(body)
	if error != nil {
		log.Panicln("Call strapi is error!.", error.Error())
		responseStruct.Error = error.Error()
		responseStruct.Status = false
		responseStruct.Payload = string(response)
		c.JSON(http.StatusBadRequest, responseStruct)
		return
	}
	responseStruct.Error = ""
	responseStruct.Status = true
	responseStruct.Payload = string(response)
	c.IndentedJSON(http.StatusOK, responseStruct)
}
