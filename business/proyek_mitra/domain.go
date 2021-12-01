package proyek_mitra

import (
	"context"

	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model
	NamaKelompokTani     string `json:"nama_kelompok_tani"`
	NamaProyek           string `json:"nama_proyek"`
	PengalamanBertani    int    `json:"pengalaman_bertani"`
	PersentaseKeuntungan int    `json:"persentase_keuntungan"`
	NominalPengajuan     int    `json:"nominal_pengajuan"`
	PeriodePanen         int    `json:"periode_panen"`
	AtasNamaRekening     string `json:"atas_nama_rekening"`
	DeskripsiProyek      string `json:"deskripsi_proyek"`
	DeskripsiUmum        string `json:"deskripsi_umum"`
	StatusMitra          bool   `json:"status_mitra"`
	LinkDokumen          string `json:"link_dokumen"`
	LinkFoto             string `json:"link_foto"`
	UserID               uint   `json:"user_id"`
	SektorId             uint   `json:"sektor_id"`
	BankId               uint   `json:"bank_id"`
	ProyekID             uint   `json:"proyek_id"`
}

type Usecase interface {
	GetAllDataProyek(ctx context.Context) ([]Domain, error)
	GetAllDataByIdController(ctx context.Context, id int) (Domain, error)
	GetAllDataBySektorController(ctx context.Context, id int) ([]Domain, error)
	CreateProyekController(ctx context.Context, data Domain) (Domain, error)
	UpdateProyekController(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteProyekController(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	GetAllDataProyek(ctx context.Context) ([]Domain, error)
	GetAllDataById(ctx context.Context, id int) (Domain, error)
	GetAllDataBySektor(ctx context.Context, id int) ([]Domain, error)
	CreateProyek(ctx context.Context, data Domain) (Domain, error)
	UpdateProyek(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteProyek(ctx context.Context, id int) (Domain, error)
}
