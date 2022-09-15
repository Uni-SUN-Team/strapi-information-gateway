package controllers

import (
	"unisun/api/unisun-strapi-inquiry/src/models"
	"unisun/api/unisun-strapi-inquiry/src/ports/services"

	"github.com/gin-gonic/gin"
	"github.com/narawichsaphimarn/handlercontrol/response/constants"
	serviceResp "github.com/narawichsaphimarn/handlercontrol/response/services"
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
	if err := c.ShouldBindJSON(&body); err != nil {
		code := constants.InternalServerError
		msg := serviceResp.HttpMessage(code)
		c.AbortWithStatusJSON(code, serviceResp.MapResponseUnsuccess(code, msg, err.Error(), nil))
	}
	response, err := srv.Service.CallStrapi(body)
	if err != nil {
		code := constants.InternalServerError
		msg := serviceResp.HttpMessage(code)
		c.AbortWithStatusJSON(code, serviceResp.MapResponseUnsuccess(code, msg, err.Error(), nil))
	}
	code := constants.OK
	msg := serviceResp.HttpMessage(code)
	c.AbortWithStatusJSON(code, serviceResp.MapResponseSuccess(code, msg, response))
}
