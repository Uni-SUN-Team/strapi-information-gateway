package src

import (
	"os"
	"unisun/api/strapi-information-gateway/src/route"

	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	r := gin.Default()
	g := r.Group(os.Getenv("CONTEXT_PATH") + "/api")
	{
		route.Services(g)
	}
	return r
}
