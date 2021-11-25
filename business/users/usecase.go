package users

import (
	"context"
	"errors"

	_middleware "investaBackend/app/middlewares"
	"time"
)

type UserUsecase struct {
	repo UserRepoInterface
	ctx  time.Duration
	jwt  *_middleware.ConfigJWT

	//usecase lain
	//repolain
}

func NewUsecase(userRepo UserRepoInterface, contextTimeout time.Duration, configJWT *_middleware.ConfigJWT) UserUseCaseInterface {
	return &UserUsecase{
		repo: userRepo,
		ctx:  contextTimeout,
		jwt:  configJWT,
	}
}

func (usecase *UserUsecase) Login(domain Domain, ctx context.Context) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, errors.New("Email Kosong")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("Password Kosong")
	}

	user, err := usecase.repo.Login(domain, ctx)
	if err != nil {
		return Domain{}, err
	}
	tokens := usecase.jwt.GenererateToken(user.Id)
	user.Token = tokens

	return user, nil
}

func (usecase *UserUsecase) GetAllUsers(ctx context.Context) ([]Domain, error) {
	return []Domain{}, nil
}
