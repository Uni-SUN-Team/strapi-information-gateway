package ports

import "unisun/api/strapi-information-gateway/src/model"

type ServiceIncome interface {
	CallStrapi(payload model.Payload) ([]byte, error)
}
