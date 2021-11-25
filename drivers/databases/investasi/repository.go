package investasi

import (
	"context"
	"investaBackend/business/investasi"

	"gorm.io/gorm"
)

type InvestasiRepository struct {
	db *gorm.DB
}

func NewInvestasiRepository(gormDB *gorm.DB) investasi.DomainInvestasiRepository {
	return &InvestasiRepository{
		db: gormDB,
	}
}

func (repo *InvestasiRepository) GetAllInvestor(domain investasi.DomainInvestasi, ctx context.Context) (, error) {
	investasiDb := FromDomain(domain)

	err := repo.db.Find()
}
