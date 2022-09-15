package controllers

import (
	"log"
	"net/http"

	"unisun/api/unisun-strapi-inquiry/src/models"
	"unisun/api/unisun-strapi-inquiry/src/ports/services"

	"github.com/gin-gonic/gin"
)

type service struct {
	Service services.ServiceStrapi
}

func NewControllerStrapi(serviceStrapi services.ServiceStrapi) *service {
	return &service{
		Service: serviceStrapi,
	}
}

func (srv *service) CallStrapi(c *gin.Context) {
	var body = models.Payload{}
	var responseStruct = models.Response{}
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
