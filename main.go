package main

import (
	"log"
	"unisun/api/unisun-strapi-inquiry/src"

	"unisun/api/unisun-strapi-inquiry/src/configs/environment"
	"unisun/api/unisun-strapi-inquiry/src/constants"
)

func main() {
	err := environment.LoadEnvironment()
	if err != nil {
		log.Fatal(err)
	}
	r := src.App()
	port := environment.ENV.App.Port
	if port == "" {
		r.Run(":" + constants.PORT)
	} else {
		r.Run(":" + port)
	}
}
