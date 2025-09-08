package provider

import (
	"context"
	"github.com/Microservices-Gopher-Hex/provider-service/domain/model"
	"github.com/Microservices-Gopher-Hex/provider-service/domain/repository"
	"github.com/google/uuid"
)

type CreateUseCase struct {
	Repo repository.ProviderRepository
}

func (uc *CreateUseCase) Do(ctx context.Context, name string, email *string) (model.Provider, error) {
	p := model.Provider{ID: uuid.NewString(), Name: name, Email: email}
	return uc.Repo.Create(ctx, p)
}
