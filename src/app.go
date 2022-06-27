package src

import (
	"unisun/api/strapi-information-gateway/docs"
	"unisun/api/strapi-information-gateway/src/controller"
	"unisun/api/strapi-information-gateway/src/route"
	"unisun/api/strapi-information-gateway/src/service"
	"unisun/api/strapi-information-gateway/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
	docs.SwaggerInfo.Version = viper.GetString("app.version")
	docs.SwaggerInfo.Host = viper.GetString("app.host")
	docs.SwaggerInfo.BasePath = viper.GetString("app.context_path") + viper.GetString("app.root_path")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	healController := controller.NewControllerHealthCheckHandler()

	utilAdap := utils.New(viper.GetString("endpoint.strapi.access_token"))
	serviceAdap := service.New(utilAdap)
	controllerAdap := controller.NewControllerServiceIncomeAdapter(serviceAdap)
	routeAdap := route.NewRouteService(controllerAdap)
	r := gin.Default()
	g := r.Group(viper.GetString("app.context_path") + viper.GetString("app.root_path") + "/v1")
	{
		g.GET("/healcheck", healController.HealthCheckHandler)
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		g.StaticFile("/license", "./LICENSE")
		routeAdap.Services(g)
	}
	return r
}
