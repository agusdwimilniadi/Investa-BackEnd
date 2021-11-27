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

// func (repo UserRepository) Login(domain users.Domain, ctx context.Context) (users.Domain, error) {
// 	userDb := FromDomain(domain)
// 	err := repo.db.Where("email = ?", userDb.Email).First(&userDb).Error
// 	hashing, err := encrypt.Hash(domain.Password)

// 	if err != nil && !encrypt.ValidateHash(domain.Password, hashing) {
// 		return users.Domain{}, err
// 	}

// 	return userDb.ToDomain(), nil
// }
func (repo UserRepository) GetAllUsers(ctx context.Context) ([]users.Domain, error) {
	return []users.Domain{}, nil
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
