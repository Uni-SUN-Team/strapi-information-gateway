package route

import (
	"unisun/api/strapi-information-gateway/controller"

	"github.com/gin-gonic/gin"
)

func Services(r *gin.RouterGroup) {
	r.POST("/strapi", controller.StrapiGateway)
}
