package controllers

import "github.com/gin-gonic/gin"

type ControllerHealthCheck interface {
	HealthCheckHandler(c *gin.Context)
}
