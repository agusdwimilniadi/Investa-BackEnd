package sektor

import "context"

type Domain struct {
	ID         uint
	NamaSektor string
	// ProyekProyek []ProyekMitra
}

type UseCase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	FindByID(id int) (Domain, error)
}
