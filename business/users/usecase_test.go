package users_test

import (
	"errors"
	middleware "investaBackend/app/middlewares"
	"investaBackend/business/proyek_mitra"
	"investaBackend/business/users"
	"investaBackend/business/users/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	userRepoInterfaceMock mocks.UserRepoInterface
	userUseCaseInterface  users.UserUseCaseInterface
	userDataDummyLogin    users.Domain
	proyekDomain          proyek_mitra.Domain
)

func setup() {
	configJWT := middleware.ConfigJWT{
		SecretJWT:       "123",
		ExpiresDuration: 2,
	}
	userUseCaseInterface = users.NewUsecase(&userRepoInterfaceMock, time.Hour*1, &configJWT)
	userDataDummyLogin = users.Domain{
		Model: gorm.Model{
			ID: 1,
		},
		Email:    "agusdwimill@gmail.com",
		Name:     "Agus",
		Password: "123",
		Token:    "",
	}
}

func TestRegister(t *testing.T) {
	setup()

	t.Run("Succes Register", func(t *testing.T) {
		userRepoInterfaceMock.On("Register", mock.AnythingOfType("int"), mock.Anything).Return(errors.New("Error")).Once()
		// usersTest, err := userUseCaseInterface.Register(context.Background(), users.Domain)

	})

	// var requestRegisterDomain = users.Domain{
	// 	Name:     "Agus",
	// 	Email:    "agusdwimill@gmail.com",
	// 	Password: "123",
	// }
	// users, err := userUseCaseInterface.Register(context.Background(), &requestRegisterDomain)
	// assert.Equal(t, usersTest, users.Domain{})

}
func TestGetByEmail() {

}
