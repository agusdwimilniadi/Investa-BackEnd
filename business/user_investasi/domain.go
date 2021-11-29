package user_investasi

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Email     string
	Name      string
	Password  string
	Token     string
}

type UserInvestasiUseCaseInterface interface {
	Login(ctx context.Context, email, password string) (string, error)
	Register(ctx context.Context, domain *Domain) error
}

type UserInvestasiRepoInterface interface {
	Register(domain *Domain, ctx context.Context) error
	GetByEmail(ctx context.Context, email string) (Domain, error)
}
