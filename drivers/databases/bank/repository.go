package bank

import (
	"context"
	"investaBackend/business/bank"
	"log"

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
	log.Println("sebelum create ke db")
	err := repo.db.Create(&bankDb).Error
	log.Println("setelah create ke db")
	// result := repo.db.Create(&bankDb)

	if err != nil {
		return bank.Domain{}, err
	}
	return bankDb.ToDomain(), nil
}
