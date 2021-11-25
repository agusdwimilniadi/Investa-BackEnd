package response

import (
	"investaBackend/business/user_investasi"

	"gorm.io/gorm"
)

type UserInvestasiResponse struct {
	gorm.Model
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func FromDomain(domain user_investasi.Domain) UserInvestasiResponse {
	return UserInvestasiResponse{
		Model: gorm.Model{
			ID:        domain.Id,
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
			DeletedAt: domain.DeletedAt,
		},
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
		Token:    domain.Token,
	}
}
