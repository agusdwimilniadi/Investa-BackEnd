package sektor

import (
	"context"
	"investaBackend/business/sektor"

	"gorm.io/gorm"
)

type SektorRepository struct {
	db *gorm.DB
}

func NewSektorRepository(gormDb *gorm.DB) sektor.Repository {
	return &SektorRepository{
		db: gormDb,
	}
}

func (repo *SektorRepository) InsertSektor(domain sektor.Domain, ctx context.Context) (sektor.Domain, error) {
	sektorDb := FromDomain(domain)
	err := repo.db.Create(&sektorDb).Error

	if err != nil {
		return sektor.Domain{}, err
	}
	return sektorDb.ToDomain(), nil
}
