package controller

import (
	"log"
	"net/http"
	"unisun/api/strapi-information-gateway/src/model"
	"unisun/api/strapi-information-gateway/src/service"

	"github.com/gin-gonic/gin"
)

func StrapiGateway(c *gin.Context) {
	var body = model.Payload{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false, "payload": nil})
		return
	}
	response, error := service.CallStrapi(body)
	if error != nil {
		log.Panicln("Call strapi is error!.", error.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": error, "status": false, "payload": response})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"error": nil, "status": true, "payload": string(response)})
}
