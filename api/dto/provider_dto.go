package dto

type CreateProviderRequest struct {
	Name  string  `json:"name" validate:"required,min=2"`
	Email *string `json:"email" validate:"omitempty,email"`
}

type ProviderResponse struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email *string `json:"email"`
}
