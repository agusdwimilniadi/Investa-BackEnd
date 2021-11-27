package investasi

import (
	"context"
	"time"
)

type InvestasiUseCase struct {
	repo Repository
	ctx  time.Duration
}

func NewInvestasiUseCase(investasiRepo Repository, contextTimeout time.Duration) Usecase {
	return &InvestasiUseCase{
		repo: investasiRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *InvestasiUseCase) InsertInvestasi(domain DomainInvestasi, ctx context.Context) (DomainInvestasi, error) {
	invest, err := usecase.repo.InsertInvestasi(domain, ctx)

	if err != nil {
		return DomainInvestasi{}, err
	}
	return invest, nil
}
