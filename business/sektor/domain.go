package sektor

import (
	"context"
)

type Domain struct {
	Id         uint
	NamaSektor string
	// ProyekBank []proyek_mitra
}

type UseCase interface {
	InsertSektor(domain Domain, ctx context.Context) (Domain, error)
}

type Repository interface {
	InsertSektor(domain Domain, ctx context.Context) (Domain, error)
}
