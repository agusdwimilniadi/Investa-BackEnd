package bank

import (
	"context"
	"errors"
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
func (usecase *BankUseCase) GetAllDataBank(ctx context.Context) ([]Domain, error) {
	banks, err := usecase.repo.GetAllDataBank(ctx)

	if err != nil {
		return []Domain{}, errors.New("Error")
	}
	return banks, nil
}
