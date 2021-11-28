package proyek_mitra

import (
	"context"
	"time"
)

type ProyekMitraUseCase struct {
	repo Repository
	ctx  time.Duration
}

func NewProyekMitraUseCase(repo Repository, timeout time.Duration) Usecase {
	return &ProyekMitraUseCase{
		repo: repo,
		ctx:  timeout,
	}
}

func (uc *ProyekMitraUseCase) GetAllDataProyek(ctx context.Context) ([]Domain, error) {
	proyek, err := uc.repo.GetAllDataProyek(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return proyek, nil
}

func (uc *ProyekMitraUseCase) GetAllDataByIdController(ctx context.Context, id int) (Domain, error) {
	proyekMitra, err := uc.repo.GetAllDataById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return proyekMitra, nil
}

func (uc *ProyekMitraUseCase) CreateProyekController(ctx context.Context, data Domain) (Domain, error) {
	proyekMitra, err := uc.repo.CreateProyek(ctx, data)
	if err != nil {
		return Domain{}, err
	}
	return proyekMitra, nil
}
