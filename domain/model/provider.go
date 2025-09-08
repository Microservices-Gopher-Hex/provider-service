package domain

import "time"

type Provider struct {
	ID        string
	Name      string
	Email     *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
