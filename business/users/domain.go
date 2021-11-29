package users

import (
	"context"

	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model
	Email    string
	Name     string
	Password string
	Token    string
	// ProyekUser []proyek_mitra.ProyekMitra
}

type UserUseCaseInterface interface {
	Login(ctx context.Context, email, password string) (string, error)
	Register(ctx context.Context, domain *Domain) error
}

type UserRepoInterface interface {
	Register(domain *Domain, ctx context.Context) error
	GetByEmail(ctx context.Context, email string) (Domain, error)
}
