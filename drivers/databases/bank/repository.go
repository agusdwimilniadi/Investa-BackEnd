package bank

import (
	"context"
	"investaBackend/business/bank"

	"gorm.io/gorm"
)

type BankRepository struct {
	db *gorm.DB
}

func NewBankRepository(gormDb *gorm.DB) bank.Repository {
	return &BankRepository{
		db: gormDb,
	}
}

func (repo *BankRepository) InsertBank(domain bank.Domain, ctx context.Context) (bank.Domain, error) {
	bankDb := FromDomain(domain)
	err := repo.db.Create(&bankDb).Error

	if err != nil {
		return bank.Domain{}, err
	}
	return bankDb.ToDomain(), nil
}
