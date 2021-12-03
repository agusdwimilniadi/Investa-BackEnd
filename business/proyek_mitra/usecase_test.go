package proyek_mitra_test

import (
	"context"
	"errors"
	"investaBackend/business/proyek_mitra"
	"investaBackend/business/proyek_mitra/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	proyekMitraMockInterface    mocks.Repository
	proyekMitraUseCaseInterface proyek_mitra.Usecase
	proyekMitraDomain           proyek_mitra.Domain
)

func setup() {
	proyekMitraUseCaseInterface = proyek_mitra.NewProyekMitraUseCase(&proyekMitraMockInterface, time.Hour*1)
	proyekMitraDomain = proyek_mitra.Domain{
		NamaKelompokTani:     "Sugeng",
		NamaProyek:           "Tani Padi",
		PengalamanBertani:    1,
		PersentaseKeuntungan: 1,
		NominalPengajuan:     1,
		PeriodePanen:         1,
		AtasNamaRekening:     "Angga",
		DeskripsiProyek:      "Desc",
		DeskripsiUmum:        "Desc",
		StatusMitra:          false,
		LinkDokumen:          "Link",
		LinkFoto:             "Link",
		UserID:               1,
		SektorId:             1,
		BankId:               1,
		ProyekID:             1,
	}

}
func TestGetAllDataProyek(t *testing.T) {
	setup()
	t.Run("Succes get data", func(t *testing.T) {
		proyekMitraMockInterface.On("GetAllDataProyek", mock.Anything).Return([]proyek_mitra.Domain{proyekMitraDomain}, nil).Once()

		domain, err := proyekMitraUseCaseInterface.GetAllDataProyek(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, "Sugeng", domain[0].NamaKelompokTani)
	})
	t.Run("Error get data", func(t *testing.T) {
		proyekMitraMockInterface.On("GetAllDataProyek", mock.Anything).Return([]proyek_mitra.Domain{}, errors.New("Error")).Once()

		domain, err := proyekMitraUseCaseInterface.GetAllDataProyek(context.Background())
		assert.Nil(t, nil, domain)
		assert.Equal(t, errors.New("Error"), err)
	})
}

func TestGetAllDataById(t *testing.T) {
	setup()
	proyekMitraMockInterface.On("GetAllDataById", mock.Anything, mock.AnythingOfType("int")).Return(proyekMitraDomain, nil).Once()
	proyekMitraMockInterface.On("GetAllDataById", mock.Anything, mock.AnythingOfType("int")).Return(proyek_mitra.Domain{}, errors.New("Error")).Once()

	t.Run("Succes get by ID", func(t *testing.T) {
		result, err := proyekMitraUseCaseInterface.GetAllDataByIdController(context.Background(), 0)
		assert.Nil(t, err)
		assert.Equal(t, proyekMitraDomain, result)
	})
	t.Run("Failed", func(t *testing.T) {
		result, err := proyekMitraUseCaseInterface.GetAllDataByIdController(context.Background(), 0)

		assert.Equal(t, errors.New("Error"), err)
		assert.Equal(t, proyek_mitra.Domain{}, result)
	})

}

func TestGetAllDataBySektor(t *testing.T) {
	setup()
	proyekMitraMockInterface.On("GetAllDataBySektor", mock.Anything, mock.AnythingOfType("int")).Return([]proyek_mitra.Domain{proyekMitraDomain}, nil).Once()
	proyekMitraMockInterface.On("GetAllDataBySektor", mock.Anything, mock.AnythingOfType("int")).Return([]proyek_mitra.Domain{}, errors.New("Error")).Once()

	t.Run("Succes get by ID", func(t *testing.T) {
		result, err := proyekMitraUseCaseInterface.GetAllDataBySektorController(context.Background(), 0)
		assert.Nil(t, err)
		assert.Equal(t, "Sugeng", result[0].NamaKelompokTani)
	})
	t.Run("Failed", func(t *testing.T) {
		result, err := proyekMitraUseCaseInterface.GetAllDataBySektorController(context.Background(), 0)

		assert.Equal(t, errors.New("Error"), err)
		assert.Nil(t, nil, result)
	})

}
