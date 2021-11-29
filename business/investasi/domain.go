package investasi

import (
	"context"

	"gorm.io/gorm"
)

type DomainInvestasi struct {
	gorm.Model
	UserInvestasiID uint
	ProyekID        uint
	Nominal         int
	LinkBuktiTf     string
	// Proyek          proyek_mitra.ProyekMitra
}

type DomainTotalInvestasi struct {
	TotalInvestasi int
}

type Usecase interface {
	InsertInvestasi(domain DomainInvestasi, ctx context.Context) (DomainInvestasi, error)
	TotalInvestasiByIdController(ctx context.Context, id int) (DomainTotalInvestasi, error)
}

type Repository interface {
	InsertInvestasi(domain DomainInvestasi, ctx context.Context) (DomainInvestasi, error)
	TotalInvestasiById(ctx context.Context, id int) (DomainTotalInvestasi, error)
}
