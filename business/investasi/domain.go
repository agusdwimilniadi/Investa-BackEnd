package investasi

import (
	"context"
	"investaBackend/drivers/databases/proyek_mitra"

	"gorm.io/gorm"
)

type DomainInvestasi struct {
	gorm.Model
	UserInvestasiID uint
	ProyekID        uint
	Nominal         int
	LinkBuktiTf     string
	Proyek          proyek_mitra.ProyekMitra
}

type Usecase interface {
	InsertInvestasi(domain DomainInvestasi, ctx context.Context) (DomainInvestasi, error)
}

type Repository interface {
	InsertInvestasi(domain DomainInvestasi, ctx context.Context) (DomainInvestasi, error)
}
