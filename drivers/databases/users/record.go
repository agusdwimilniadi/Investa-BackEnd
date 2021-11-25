package users

import (
	"investaBackend/business/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func (user User) ToDomain() users.Domain {
	return users.Domain{
		Id:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		Email:     user.Email,
		Name:      user.Name,
		Password:  user.Password,
	}

}

func FromDomain(domain users.Domain) User {
	return User{
		Model: gorm.Model{
			ID:        domain.Id,
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
			DeletedAt: domain.DeletedAt,
		},
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
	}
}
