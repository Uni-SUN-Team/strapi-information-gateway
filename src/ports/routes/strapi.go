package routes

import "github.com/gin-gonic/gin"

type RouteStrapi interface {
	Services(r *gin.RouterGroup)
}
