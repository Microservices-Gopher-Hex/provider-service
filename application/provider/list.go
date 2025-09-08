package provider

import (
	"context"
	"github.com/Microservices-Gopher-Hex/provider-service/domain/model"
	"github.com/Microservices-Gopher-Hex/provider-service/domain/repository"
)

type ListUseCase struct {
	Repo repository.ProviderRepository
}

func (uc *ListUseCase) Do(ctx context.Context, limit, offset int32) ([]model.Provider, error) {
	return uc.Repo.List(ctx, limit, offset)
}
