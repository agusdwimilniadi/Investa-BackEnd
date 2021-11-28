package proyek_mitra

import (
	"context"
	"investaBackend/business/proyek_mitra"

	"gorm.io/gorm"
)

type ProyekMitraRepository struct {
	db *gorm.DB
}

func NewProyekMitraRepository(gormDB *gorm.DB) proyek_mitra.Repository {
	return &ProyekMitraRepository{
		db: gormDB,
	}
}

func (db *ProyekMitraRepository) GetAllDataProyek(ctx context.Context) ([]proyek_mitra.Domain, error) {
	var currentProyek []ProyekMitra
	result := db.db.Find(&currentProyek)

	if result.Error != nil {
		return []proyek_mitra.Domain{}, result.Error
	}
	return ToDomains(currentProyek), nil
}

func (db *ProyekMitraRepository) GetAllDataById(ctx context.Context, id int) (proyek_mitra.Domain, error) {
	var currentProyek ProyekMitra
	result := db.db.Where("id = ?", id).First(&currentProyek)
	if result.Error != nil {
		return proyek_mitra.Domain{}, result.Error
	}
	return currentProyek.ToDomain(), nil
}

func (db *ProyekMitraRepository) CreateProyek(ctx context.Context, data proyek_mitra.Domain) (proyek_mitra.Domain, error) {
	insertProyek := FromDomain(data)
	result := db.db.Create(&insertProyek)

	if result.Error != nil {
		return proyek_mitra.Domain{}, result.Error
	}
	return insertProyek.ToDomain(), nil
}
