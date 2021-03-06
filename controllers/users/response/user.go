package response

import (
	"investaBackend/business/users"
	"time"

	"gorm.io/gorm"
)

type UserResponse struct {
	Id        uint           `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	Email     string         `json:"email"`
	Name      string         `json:"name"`
	Password  string         `json:"password"`
	Token     string         `json:"token"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		Id:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  domain.Password,
		Token:     domain.Token,
	}
}
