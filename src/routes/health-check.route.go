package routes

import (
	"unisun/api/unisun-strapi-inquiry/src/ports/controllers"

	"github.com/gin-gonic/gin"
)

type controllerHealthCheck struct {
	Controller controllers.ControllerHealthCheck
}

func NewHealthCheck(_controllerHealthCheck controllers.ControllerHealthCheck) *controllerHealthCheck {
	return &controllerHealthCheck{
		Controller: _controllerHealthCheck,
	}
}

func (srv *controllerHealthCheck) HealthCheckRoute(g *gin.RouterGroup) {
	g.GET("/health-check", srv.Controller.HealthCheckHandler)
}
