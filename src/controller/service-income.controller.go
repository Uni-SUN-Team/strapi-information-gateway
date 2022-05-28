package controller

import (
	"net/http"
	"unisun/api/strapi-information-gateway/src/logging"
	"unisun/api/strapi-information-gateway/src/model"
	"unisun/api/strapi-information-gateway/src/service"

	"github.com/gin-gonic/gin"
)

// StrapiInformationGatewayHandler godoc
// @summary      Strapi Information Gateway
// @description  Strapi Information Gateway for the service
// @id           StrapiInformationGatewayHandler
// @produce      json
// @success     200  {object}  model.Response  "OK"
// @failure		400  {object}  model.Response  "FAIL"
// @router       /strapi [post]
func StrapiGateway(c *gin.Context) {
	var body = model.Payload{}
	var responseStruct = model.Response{}
	if err := c.ShouldBindJSON(&body); err != nil {
		logging.Println("Bind json body is error.", err.Error())
		responseStruct.Error = err.Error()
		responseStruct.Status = false
		responseStruct.Payload = ""
		c.JSON(http.StatusBadRequest, responseStruct)
		return
	}
	response, error := service.CallStrapi(body)
	if error != nil {
		logging.Println("Call strapi is error!.", error.Error())
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
