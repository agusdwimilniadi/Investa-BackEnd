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

func (repo *InvestasiRepository) TotalInvestasiById(ctx context.Context, id int) (investasi.DomainTotalInvestasi, error) {
	var total TotalInvestasi
	// result := repo.db.Raw("SELECT SUM(nominal) FROM investasis WHERE proyek_id = ?", id).First(&total).Error
	results := repo.db.Table("investasis").Where("proyek_id = ?", id).Select("sum(nominal)").Row()
	result := results.Scan(&total)
	if result != nil {
		return investasi.DomainTotalInvestasi{}, result
	}
	return total.ToDomainTotal(), nil
}
