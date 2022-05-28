package config

import (
	"os"
	"unisun/api/strapi-information-gateway/src/constants"
)

func ConfigENV() {
	os.Setenv(constants.NODE, "development")
	os.Setenv(constants.PORT, "8082")
	os.Setenv(constants.CONTEXT_PATH, "/strapi-information-gateway")
	os.Setenv(constants.HOST_STRAPI, "http://localhost:1337")
	os.Setenv(constants.TOKEN_STRAPI, "6102511ec9d936311659e59bd511783c8a1f3b4b00f47b27fbc16df3a4d2594924189cfa0412d435966fc188b6573e7a378323537e3d67b3cdbe87a0791012aa8425bd11f70266afa897ec5cb1b82dbb57529da17f273d94ec336b93c0a5cf9b056f1d52ba6d9576685227508c4736f8d06a9158d37ea665fd07f5a81bb6e72e")
	os.Setenv(constants.LOG_PATH, "/Users/ns/Documents/UniSUN/strapi-information-gateway/tmp/app_log")
}
