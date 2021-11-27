package user_investasi

import (
	"context"
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

func (repo UserInvestasiRepository) Login(domain user_investasi.Domain, ctx context.Context) (user_investasi.Domain, error) {
	userDb := FromDomain(domain)
	err := repo.db.Where("email = ? AND password = ?", userDb.Email, userDb.Password).First(&userDb).Error

	if err != nil {
		return user_investasi.Domain{}, err
	}
	return userDb.ToDomain(), nil
}
