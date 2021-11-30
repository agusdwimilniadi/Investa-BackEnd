package sektor

import (
	"context"
	"time"
)

type SektorUseCase struct {
	repo Repository
	ctx  time.Duration
}

func NewSektorUseCase(sektorRepo Repository, contextTime time.Duration) UseCase {
	return &SektorUseCase{
		repo: sektorRepo,
		ctx:  contextTime,
	}
}

func (usecase *SektorUseCase) InsertSektor(domain Domain, ctx context.Context) (Domain, error) {
	sektors, err := usecase.repo.InsertSektor(domain, ctx)

	if err != nil {
		return Domain{}, err
	}
	return sektors, nil
}
