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
	// Proyek          ProyekMitra
}

type DomainInvestasiUseCase interface {
	GetAllInvestor(domain DomainInvestasi, ctx context.Context) (DomainInvestasi, error)
}

type DomainInvestasiRepository interface {
	GetAllInvestor(domain DomainInvestasi, ctx context.Context) (DomainInvestasi, error)
}
