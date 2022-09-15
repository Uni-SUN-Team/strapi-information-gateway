package controllers

import "github.com/gin-gonic/gin"

type ControllerStrapi interface {
	CallStrapi(c *gin.Context)
}
