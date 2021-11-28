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
	// ProyekUser []proyek_mitra.ProyekMitra `gorm:"foreignKey:UserId"`
}

func (user *User) ToDomain() users.Domain {
	return users.Domain{
		Model: gorm.Model{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		// ProyekUser: user.ProyekUser,
	}

}

func FromDomain(domain users.Domain) *User {
	return &User{
		Model: gorm.Model{
			ID:        domain.ID,
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
			DeletedAt: domain.DeletedAt,
		},
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
	}
}
