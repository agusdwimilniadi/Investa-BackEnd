package sektor_test

import (
	"context"
	"investaBackend/business/sektor"
	"investaBackend/business/sektor/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	sektorRepoInterfaceMock mocks.Repository
	sektorUseCaseInterface  sektor.UseCase
	sektorDataDummyInsert   sektor.Domain
)

func setup() {
	sektorUseCaseInterface = sektor.NewSektorUseCase(&sektorRepoInterfaceMock, time.Hour*1)
	sektorDataDummyInsert = sektor.Domain{
		Id:         1,
		NamaSektor: "Pertanian",
	}
}

func TestInsertSektor(t *testing.T) {
	setup()
	t.Run("Succes insert", func(t *testing.T) {
		sektorRepoInterfaceMock.On("InsertSektor", mock.AnythingOfType("sektor.Domain"), mock.Anything).Return(sektorDataDummyInsert, nil).Once()

		var requestInsertSektorDomain = sektor.Domain{
			NamaSektor: "Pertanian",
		}
		domain, err := sektorUseCaseInterface.InsertSektor(requestInsertSektorDomain, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, sektorDataDummyInsert, domain)

	})
}
