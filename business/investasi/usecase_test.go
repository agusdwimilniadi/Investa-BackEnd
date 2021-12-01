package investasi_test

import (
	"context"
	"investaBackend/business/investasi"
	"investaBackend/business/investasi/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	investasiRepoInterfaceMock mocks.Repository
	investasiUseCaseInterface  investasi.Usecase
	investasiDataDummyInsert   investasi.DomainInvestasi
	investasiDataDummyTotal    investasi.DomainTotalInvestasi
)

func setup() {
	investasiUseCaseInterface = investasi.NewInvestasiUseCase(&investasiRepoInterfaceMock, time.Hour*1)
	investasiDataDummyInsert = investasi.DomainInvestasi{
		UserInvestasiID: 1,
		ProyekMitraID:   1,
		Nominal:         10,
		LinkBuktiTf:     "link",
	}
	investasiDataDummyTotal = investasi.DomainTotalInvestasi{
		TotalInvestasi: 10000,
	}
}

func TestInsertInvestasi(t *testing.T) {
	setup()
	t.Run("Succes insert", func(t *testing.T) {
		investasiRepoInterfaceMock.On("InsertInvestasi", mock.AnythingOfType("investasi.DomainInvestasi"), mock.Anything).Return(investasiDataDummyInsert, nil).Once()

		var requestInsertInvestasiDomain = investasi.DomainInvestasi{
			UserInvestasiID: 1,
			ProyekMitraID:   1,
			Nominal:         10,
			LinkBuktiTf:     "link",
		}
		domain, err := investasiUseCaseInterface.InsertInvestasi(requestInsertInvestasiDomain, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, investasiDataDummyInsert, domain)
	})
}

func TestTotalTotalInvestasiByIdController(t *testing.T) {
	setup()
	t.Run("Succes get total", func(t *testing.T) {
		investasiRepoInterfaceMock.On("TotalInvestasiById", mock.Anything, mock.AnythingOfType("int")).Return(investasiDataDummyTotal, nil).Once()

		domain, err := investasiUseCaseInterface.TotalInvestasiByIdController(context.Background(), 1)
		assert.Equal(t, nil, err)
		assert.Equal(t, investasiDataDummyTotal, domain)
	})
}
