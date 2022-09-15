package services

import "unisun/api/unisun-strapi-inquiry/src/models"

type ServiceStrapi interface {
	CallStrapi(payload models.Payload) ([]byte, error)
}
