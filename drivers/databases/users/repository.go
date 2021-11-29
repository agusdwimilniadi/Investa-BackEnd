package users

import (
	"context"
	"errors"
	"investaBackend/business/users"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(gormDb *gorm.DB) users.UserRepoInterface {
	return &UserRepository{
		db: gormDb,
	}
}

func (repo *UserRepository) GetByEmail(ctx context.Context, emailCheck string) (users.Domain, error) {
	rec := User{}
	err := repo.db.Where("email = ?", emailCheck).First(&rec).Error
	if err != nil {
		return users.Domain{}, errors.New("email atau password salah")
	}
	return rec.ToDomain(), nil
}

func (repo *UserRepository) Register(user *users.Domain, ctx context.Context) error {
	record := FromDomain(*user)
	result := repo.db.Create(record).Error

	if result != nil {
		return result
	}
	return nil
}
