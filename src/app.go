package src

import (
	"os"
	"unisun/api/strapi-information-gateway/docs"
	"unisun/api/strapi-information-gateway/src/constants"
	"unisun/api/strapi-information-gateway/src/controller"
	"unisun/api/strapi-information-gateway/src/route"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @termsOfService 		http://swagger.io/terms/
// @contact.name 		API Support
// @contact.url 		http://www.swagger.io/support
// @contact.email 		support@swagger.io

// @license.name 		MIT License Copyright (c) 2022 Uni-SUN-Team
// @license.url 		https://api.unisun.dynu.com/strapi-information-gateway/api/license

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func App() *gin.Engine {
	docs.SwaggerInfo.Title = "STRAPI INFORMATION GATEWAY API"
	docs.SwaggerInfo.Description = "This is a server celler to strapi server."
	docs.SwaggerInfo.Version = os.Getenv(constants.VERSION)
	docs.SwaggerInfo.Host = os.Getenv(constants.HOST)
	docs.SwaggerInfo.BasePath = os.Getenv(constants.CONTEXT_PATH) + "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()
	g := r.Group(os.Getenv(constants.CONTEXT_PATH) + "/api")
	{
		g.GET("/healcheck", controller.HealthCheckHandler)
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		g.StaticFile("/license", "./LICENSE")
		route.Services(g)
	}
	return r
}
