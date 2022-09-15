package src

import (
	"strings"
	"unisun/api/unisun-strapi-inquiry/docs"

	"unisun/api/unisun-strapi-inquiry/src/configs/environment"
	"unisun/api/unisun-strapi-inquiry/src/constants"
	"unisun/api/unisun-strapi-inquiry/src/controllers"
	"unisun/api/unisun-strapi-inquiry/src/routes"
	"unisun/api/unisun-strapi-inquiry/src/services"

	"github.com/gin-gonic/gin"
	"github.com/narawichsaphimarn/handlercontrol"
	"github.com/narawichsaphimarn/handlercontrol/models"
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
	appEnv := environment.ENV.App
	ginEnv := environment.ENV.Gin
	swagEnv := environment.ENV.Swag
	docs.SwaggerInfo.Title = swagEnv.Title
	docs.SwaggerInfo.Description = swagEnv.Description
	docs.SwaggerInfo.Version = swagEnv.Version
	docs.SwaggerInfo.Host = swagEnv.Host
	docs.SwaggerInfo.BasePath = strings.Join([]string{appEnv.ContextPath, ginEnv.RootPath, ginEnv.Version}, "/")
	docs.SwaggerInfo.Schemes = strings.Split(swagEnv.Schemes, ",")

	r := gin.Default()
	r.SetTrustedProxies(strings.Split(ginEnv.Configs.TrustedProxies, ","))
	g := r.Group(strings.Join([]string{appEnv.ContextPath, ginEnv.RootPath, ginEnv.Version}, "/"))
	{
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		g.StaticFile("/license", "./LICENSE")
		mapRouteHealthCheck(g)
		mapRouteStrapi(g)
	}
	return r
}

func mapRouteHealthCheck(g *gin.RouterGroup) {
	controller := controllers.NewControllerHealthCheckHandler()
	route := routes.NewHealthCheck(controller)
	route.HealthCheckRoute(g)
}

func mapRouteStrapi(g *gin.RouterGroup) {
	endEnv := environment.ENV.Endpoint
	httpForm := models.HttpRequestForm{}
	httpForm.TimeOut = 60000
	httpForm.ContentType = constants.CONTENT_TYPE
	httpForm.Authorization = struct {
		Token string
		Type  string
	}{
		Token: endEnv.Strapi.Token,
	}
	http := handlercontrol.NewHttpRequest(httpForm)
	service := services.New(http)
	controller := controllers.NewControllerStrapi(service)
	route := routes.NewStrapi(controller)
	route.StapiRoute(g)
}
