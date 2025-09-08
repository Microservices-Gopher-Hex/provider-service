package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/Microservices-Gopher-Hex/provider-service/adapter/repository/postgres/entity"
	"github.com/Microservices-Gopher-Hex/provider-service/adapter/repository/postgres/mapper"
	"github.com/Microservices-Gopher-Hex/provider-service/domain/model"
	"github.com/Microservices-Gopher-Hex/provider-service/domain/repository"
	"gorm.io/gorm"
)

type ProviderRepository struct{ db *gorm.DB }

func New(db *gorm.DB) repository.ProviderRepository { return &ProviderRepository{db: db} }

func (r *ProviderRepository) Create(ctx context.Context, p model.Provider) (model.Provider, error) {
	m := mapper.ToEntity(p)
	now := time.Now()
	m.CreatedAt, m.UpdatedAt = now, now
	if err := r.db.WithContext(ctx).Create(&m).Error; err != nil {
		return model.Provider{}, err
	}
	return mapper.ToDomain(m), nil
}

func (r *ProviderRepository) Get(ctx context.Context, id string) (model.Provider, error) {
	var m entity.Provider
	if err := r.db.WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Provider{}, err
		}
		return model.Provider{}, err
	}
	return mapper.ToDomain(m), nil
}

func (r *ProviderRepository) List(ctx context.Context, limit, offset int32) ([]model.Provider, error) {
	if limit <= 0 {
		limit = 50
	}
	var rows []entity.Provider
	if err := r.db.WithContext(ctx).Order("created_at DESC").Limit(int(limit)).Offset(int(offset)).Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]model.Provider, 0, len(rows))
	for _, m := range rows {
		out = append(out, mapper.ToDomain(m))
	}
	return out, nil
}
