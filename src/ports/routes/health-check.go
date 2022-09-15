package routes

import "github.com/gin-gonic/gin"

type RouteHealthCheck interface {
	HealthCheckRoute(g *gin.RouterGroup)
}
