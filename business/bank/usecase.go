package bank

import (
	"context"
	"time"
)

type BankUseCase struct {
	repo Repository
	ctx  time.Duration
}

func NewBankUseCase(bankRepo Repository, contextTimeout time.Duration) UseCase {
	return &BankUseCase{
		repo: bankRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *BankUseCase) InsertBank(domain Domain, ctx context.Context) (Domain, error) {
	banks, err := usecase.repo.InsertBank(domain, ctx)

	if err != nil {
		return Domain{}, err
	}
	return banks, nil
}
