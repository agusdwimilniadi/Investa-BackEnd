package bank

import (
	"context"
	"investaBackend/business/proyek_mitra"
)

type Domain struct {
	Id         uint
	NamaBank   string
	KodeBank   string
	ProyekBank []proyek_mitra.Domain
}

type UseCase interface {
	InsertBank(domain Domain, ctx context.Context) (Domain, error)
	GetAllDataBank(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	InsertBank(domain Domain, ctx context.Context) (Domain, error)
	GetAllDataBank(ctx context.Context) ([]Domain, error)
}
