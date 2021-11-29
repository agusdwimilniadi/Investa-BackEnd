package user_investasi

import (
	"context"
	"errors"
	"investaBackend/business/user_investasi"

	"gorm.io/gorm"
)

type UserInvestasiRepository struct {
	db *gorm.DB
}

func NewUserInvestasiRepository(gormDb *gorm.DB) user_investasi.UserInvestasiRepoInterface {
	return &UserInvestasiRepository{
		db: gormDb,
	}
}

func (repo *UserInvestasiRepository) GetByEmail(ctx context.Context, emailCheck string) (user_investasi.Domain, error) {
	rec := UserInvestasi{}
	err := repo.db.Where("email = ?", emailCheck).First(&rec).Error

	if err != nil {
		return user_investasi.Domain{}, errors.New("email atau password salah")
	}
	return rec.ToDomain(), nil
}

func (repo *UserInvestasiRepository) Register(user *user_investasi.Domain, ctx context.Context) error {
	record := FromDomain(*user)
	result := repo.db.Create(record).Error

	if result != nil {
		return result
	}
	return nil
}
