package proyek_mitra_test

import (
	"investaBackend/business/proyek_mitra"
	"investaBackend/business/proyek_mitra/mocks"
	"testing"
	"time"
)

var (
	proyekMitraMockInterface    mocks.Repository
	proyekMitraUseCaseInterface proyek_mitra.Usecase
	dataInputDummy              proyek_mitra.Domain
)

func setup() {
	proyekMitraUseCaseInterface = proyek_mitra.NewProyekMitraUseCase(&proyekMitraMockInterface, time.Hour*1)
	dataInputDummy = proyek_mitra.Domain{
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
	// setup()
	// t.Run("Succes get data", func(t *testing.T) {
	// 	proyekMitraMockInterface.On("GetAllDataProyek", mock.Anything).Return(dataInputDummy, nil).Once()

	// 	domain, err := proyekMitraUseCaseInterface.GetAllDataProyek(context.Background())

	// 	assert.Equal(t, nil, err)
	// 	assert.Equal(t, dataInputDummy, domain)
	// })
}
