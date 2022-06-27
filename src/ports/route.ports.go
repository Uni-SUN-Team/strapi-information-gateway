package ports

import "github.com/gin-gonic/gin"

type RouteServiceIncome interface {
	Services(r *gin.RouterGroup)
}
