package entity

import "time"

type Provider struct {
	ID        string    `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"type:text;not null"`
	Email     *string   `gorm:"type:text"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `gorm:"not null;default:now()"`
}

func (Provider) TableName() string { return "providers" }
