package user_investasi

import (
	"context"
	"errors"
	_middleware "investaBackend/app/middlewares"
	"investaBackend/helpers/encrypt"
	"reflect"
	"time"
)

type UserInvestasiUsecase struct {
	repo UserInvestasiRepoInterface
	ctx  time.Duration
	jwt  *_middleware.ConfigJWT
}

func NewUsecase(userRepo UserInvestasiRepoInterface, contextTimeout time.Duration, configJWT *_middleware.ConfigJWT) UserInvestasiUseCaseInterface {
	return &UserInvestasiUsecase{
		repo: userRepo,
		ctx:  contextTimeout,
		jwt:  configJWT,
	}
}

func (usecase *UserInvestasiUsecase) Login(ctx context.Context, emails, password string) (string, error) {
	if emails == "" {
		return "", errors.New("email kosong")
	}
	if password == "" {
		return "", errors.New("password kosong")
	}

	user, err := usecase.repo.GetByEmail(ctx, emails)

	if err != nil {
		return "", err
	}
	if !encrypt.ValidateHash(password, user.Password) {
		return "", errors.New("email atau password salah")
	}
	tokens := usecase.jwt.GenererateToken(user.Id)

	return tokens, nil
}

func (uc *UserInvestasiUsecase) Register(ctx context.Context, userDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.ctx)
	defer cancel()

	existedUser, _ := uc.repo.GetByEmail(ctx, userDomain.Email)
	if !reflect.DeepEqual(existedUser, Domain{}) {
		return errors.New("email sudah digunakan, cari email lain")
	}
	var err error
	userDomain.Password, err = encrypt.Hash(userDomain.Password)

	if err != nil {
		return errors.New("kesalahan encrypt password")
	}
	err2 := uc.repo.Register(userDomain, ctx)

	if err2 != nil {
		return err2
	}

	return nil
}
