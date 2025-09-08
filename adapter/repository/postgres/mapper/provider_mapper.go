package mapper

import (
	"github.com/Microservices-Gopher-Hex/provider-service/adapter/repository/postgres/entity"
	"github.com/Microservices-Gopher-Hex/provider-service/domain/model"
)

func ToDomain(m entity.Provider) model.Provider {
	return model.Provider{ID: m.ID, Name: m.Name, Email: m.Email, CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt}
}

func ToEntity(p model.Provider) entity.Provider {
	return entity.Provider{ID: p.ID, Name: p.Name, Email: p.Email, CreatedAt: p.CreatedAt, UpdatedAt: p.UpdatedAt}
}
