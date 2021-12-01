package users_test

import (
	"investaBackend/business/users"
	"investaBackend/business/users/mocks"
	"testing"
	"time"

	"gorm.io/gorm"
)

var (
	userRepoInterfaceMock mocks.UserRepoInterface
	userUseCaseInterface  users.UserUseCaseInterface
	userDataDummyLogin    users.Domain
)

func setup() {
	userUseCaseInterface = users.NewUsecase(&userRepoInterfaceMock, time.Hour*1, nil)
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
	// setup()

	// t.Run("Succes Register", func(t *testing.T) {
	// 	userRepoInterfaceMock.On("Register", mock.AnythingOfType("*users.Domain"), mock.Anything).Return(nil).Once()
	// })

	// var requestRegisterDomain = users.Domain{
	// 	Name:     "Agus",
	// 	Email:    "agusdwimill@gmail.com",
	// 	Password: "123",
	// }
	// users, err = userUseCaseInterface.Register(context.Background(), &requestRegisterDomain)

}
