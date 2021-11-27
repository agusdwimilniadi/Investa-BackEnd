package investasi

import (
	"context"
	"investaBackend/business/investasi"

	"gorm.io/gorm"
)

type InvestasiRepository struct {
	db *gorm.DB
}

func NewInvestasiRepository(gormDB *gorm.DB) investasi.Repository {
	return &InvestasiRepository{
		db: gormDB,
	}
}

func (repo *InvestasiRepository) InsertInvestasi(domain investasi.DomainInvestasi, ctx context.Context) (investasi.DomainInvestasi, error) {
	investasiDb := FromDomain(domain)
	err := repo.db.Create(&investasiDb).Error

	if err != nil {
		return investasi.DomainInvestasi{}, err
	}
	return investasiDb.ToDomain(), nil
}
