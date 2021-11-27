package users

import (
	"context"
	"investaBackend/drivers/databases/proyek_mitra"

	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model
	Email      string
	Name       string
	Password   string
	Token      string
	ProyekUser []proyek_mitra.ProyekMitra
}

type UserUseCaseInterface interface {
	Login(ctx context.Context, email, password string) (string, error)
	GetAllUsers(ctx context.Context) ([]Domain, error)
	Register(ctx context.Context, domain *Domain) error
}

type UserRepoInterface interface {
	GetAllUsers(ctx context.Context) ([]Domain, error)
	Register(domain *Domain, ctx context.Context) error
	GetByEmail(ctx context.Context, email string) (Domain, error)
}
