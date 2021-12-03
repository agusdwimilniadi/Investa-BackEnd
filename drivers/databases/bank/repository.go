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

func (db *BankRepository) GetAllDataBank(ctx context.Context) ([]bank.Domain, error) {
	var currentProyek []Bank
	result := db.db.Preload("ProyekBank").Find(&currentProyek)
	if result.Error != nil {
		return []bank.Domain{}, result.Error
	}
	var domainResult []bank.Domain
	for _, val := range currentProyek {
		domainResult = append(domainResult, val.ToDomain())
	}
	return domainResult, nil
	// result := db.db.Preload("ProyekMitra").Find(&currentProyek)

	// if result.Error != nil {
	// 	return []bank.Domain{}, result.Error
	// }
	// return currentProyek.ToDomainList(), nil
}

func (repo *BankRepository) InsertBank(domain bank.Domain, ctx context.Context) (bank.Domain, error) {
	bankDb := FromDomain(domain)
	err := repo.db.Create(&bankDb).Error

	if err != nil {
		return bank.Domain{}, err
	}
	return bankDb.ToDomain(), nil
}
