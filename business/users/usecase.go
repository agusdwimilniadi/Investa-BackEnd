package users

import (
	"context"
	"errors"
	"reflect"

	_middleware "investaBackend/app/middlewares"
	"investaBackend/helpers/encrypt"
	"time"
)

type UserUsecase struct {
	repo UserRepoInterface
	ctx  time.Duration
	jwt  *_middleware.ConfigJWT
}

func NewUsecase(userRepo UserRepoInterface, contextTimeout time.Duration, configJWT *_middleware.ConfigJWT) UserUseCaseInterface {
	return &UserUsecase{
		repo: userRepo,
		ctx:  contextTimeout,
		jwt:  configJWT,
	}
}

func (usecase *UserUsecase) Login(ctx context.Context, emails, password string) (string, error) {
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
	// hash, err := encrypt.Hash(domain.Password)
	if !encrypt.ValidateHash(password, user.Password) {
		return "", errors.New("email atau password salah")
	}
	tokens := usecase.jwt.GenererateToken(user.ID)

	return tokens, nil
}

func (usecase *UserUsecase) GetAllUsers(ctx context.Context) ([]Domain, error) {
	return []Domain{}, nil
}

func (uc *UserUsecase) Register(ctx context.Context, userDomain *Domain) error {
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
