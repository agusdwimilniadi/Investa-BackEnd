package bank

import (
	"context"
)

type Domain struct {
	Id       uint
	NamaBank string
	KodeBank string
	// ProyekBank []proyek_mitra
}

type UseCase interface {
	InsertBank(domain Domain, ctx context.Context) (Domain, error)
}

type Repository interface {
	InsertBank(domain Domain, ctx context.Context) (Domain, error)
}
