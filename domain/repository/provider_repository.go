package repository

import (
	"context"
	"github.com/Microservices-Gopher-Hex/provider-service/domain/model"
)

type ProviderRepository interface {
	Create(ctx context.Context, p model.Provider) (model.Provider, error)
	Get(ctx context.Context, id string) (model.Provider, error)
	List(ctx context.Context, limit, offset int32) ([]model.Provider, error)
}
