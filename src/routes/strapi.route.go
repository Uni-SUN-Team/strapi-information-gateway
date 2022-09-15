package routes

import (
	"unisun/api/unisun-strapi-inquiry/src/ports/controllers"

	"github.com/gin-gonic/gin"
)

type controllerStrapi struct {
	Controller controllers.ControllerStrapi
}

func NewStrapi(_controllerStrapi controllers.ControllerStrapi) *controllerStrapi {
	return &controllerStrapi{
		Controller: _controllerStrapi,
	}
}

func (srv *controllerStrapi) StapiRoute(r *gin.RouterGroup) {
	r.POST("/strapi", srv.Controller.CallStrapi)
}
