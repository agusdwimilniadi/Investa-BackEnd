package user_investasi

import (
	"context"
	"errors"
	_middleware "investaBackend/app/middlewares"
	"time"
)

type UserInvestasiUsecase struct {
	repo UserInvestasiRepoInterface
	ctx  time.Duration
	jwt  *_middleware.ConfigJWT

	//usecase lain
	//repolain
}

func NewUsecase(userRepo UserInvestasiRepoInterface, contextTimeout time.Duration, configJWT *_middleware.ConfigJWT) UserInvestasiUseCaseInterface {
	return &UserInvestasiUsecase{
		repo: userRepo,
		ctx:  contextTimeout,
		jwt:  configJWT,
	}
}

func (usecase *UserInvestasiUsecase) Login(domain Domain, ctx context.Context) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, errors.New("Email kosong")
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
