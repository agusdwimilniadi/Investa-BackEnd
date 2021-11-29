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
	// var total TotalInvestasi
	// result := repo.db.Raw("SELECT SUM(nominal) FROM investasis WHERE proyek_id = ?", id).First(&total).Error
	var totalUang TotalInvestasi
	results := repo.db.Model(&Investasi{}).Select("sum(nominal) as total").Where("proyek_id = ?", id).Find(&totalUang)
	// result := results
	if results.Error != nil {
		return investasi.DomainTotalInvestasi{}, results.Error
	}
	// fmt.Println(totalUang)
	return totalUang.ToDomainTotal(), nil
}
