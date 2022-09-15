package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/narawichsaphimarn/handlercontrol/response/constants"
	"github.com/narawichsaphimarn/handlercontrol/response/services"
)

type ControllerHealthCheckHandler struct {
}

func NewControllerHealthCheckHandler() *ControllerHealthCheckHandler {
	return &ControllerHealthCheckHandler{}
}

func (*ControllerHealthCheckHandler) HealthCheckHandler(c *gin.Context) {
	var code = constants.OK
	var massage = services.HttpMessage(code)
	c.AbortWithStatusJSON(code, services.MapResponseSuccess(code, massage, nil))
}
