// package bank_test

// import (
// 	"context"
// 	"investaBackend/business/bank"
// 	"investaBackend/business/bank/mocks"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// var (
// 	bankRepoInterfaceMock mocks.Repository
// 	bankUseCaseInterface  bank.UseCase
// 	bankDataDummyInsert   bank.Domain
// )

// func setup() {
// 	bankUseCaseInterface = bank.NewBankUseCase(&bankRepoInterfaceMock, time.Hour*1)
// 	bankDataDummyInsert = bank.Domain{
// 		Id:       1,
// 		NamaBank: "BANK BNI",
// 		KodeBank: "009",
// 	}
// }

// func TestInsertBank(t *testing.T) {
// 	setup()
// 	t.Run("Succes insert", func(t *testing.T) {
// 		bankRepoInterfaceMock.On("InsertBank", mock.AnythingOfType("bank.Domain"), mock.Anything).Return(bankDataDummyInsert, nil).Once()

// 		var requestInsertBankDomain = bank.Domain{
// 			NamaBank: "BANK BNI",
// 			KodeBank: "009",
// 		}
// 		domain, err := bankUseCaseInterface.InsertBank(requestInsertBankDomain, context.Background())

// 		assert.Equal(t, nil, err)
// 		assert.Equal(t, bankDataDummyInsert, domain)
// 	})
// }
package bank_test
