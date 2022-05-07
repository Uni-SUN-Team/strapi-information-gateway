package main

import (
	"os"
	"unisun/api/strapi-information-gateway/src"
	"unisun/api/strapi-information-gateway/src/config"
)

func main() {
	if os.Getenv("PORT") != "production" {
		config.ConfigENV()
	}
	r := src.App()
	port := os.Getenv("PORT")
	if port == "" {
		r.Run(":8080")
	} else {
		r.Run(":" + port)
	}
}
