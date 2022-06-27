package route

import (
	"unisun/api/strapi-information-gateway/src/ports"

	"github.com/gin-gonic/gin"
)

type RouteService struct {
	Controller ports.ControllerServiceIncome
}

func NewRouteService(controller ports.ControllerServiceIncome) *RouteService {
	return &RouteService{
		Controller: controller,
	}
}

func (srv *RouteService) Services(r *gin.RouterGroup) {
	r.POST("/strapi", srv.Controller.StrapiGateway)
}
