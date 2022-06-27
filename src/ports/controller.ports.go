package ports

import "github.com/gin-gonic/gin"

type ControllerServiceIncome interface {
	StrapiGateway(c *gin.Context)
}

type ControllerHealCheck interface {
	HealthCheckHandler(c *gin.Context)
}
