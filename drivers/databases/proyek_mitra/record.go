package proyek_mitra

import (
	"investaBackend/drivers/databases/bank"
	"investaBackend/drivers/databases/investasi"
	"investaBackend/drivers/databases/sektor"

	"time"

	"gorm.io/gorm"
)

func (ProyekMitra) TableName() string {
	return "proyek_mitra"
}

type ProyekMitra struct {
	ID                   uint           `gorm:"primaryKey" json:"id"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	NamaKelompokTani     string         `json:"nama_kelompok_tani"`
	NamaProyek           string         `json:"nama_proyek"`
	PengalamanBertani    int            `json:"pengalaman_bertani"`
	PersentaseKeuntungan int            `json:"persentase_keuntungan"`
	NominalPengajuan     int            `json:"nominal_pengajuan"`
	PeriodePanen         int            `json:"periode_panen"`
	AtasNamaRekening     string         `json:"atas_nama_rekening"`
	DeskripsiProyek      string         `json:"deskripsi_proyek"`
	DeskripsiUmum        string         `json:"deskripsi_umum"`
	StatusMitra          bool           `json:"status_mitra"`
	LinkDokumen          string         `json:"link_dokumen"`
	LinkFoto             string         `json:"link_foto"`
	UserID               uint           `json:"user_id"`
	SektorId             uint           `json:"sektor_id"`
	Sektor               sektor.Sektor  `gorm:"foreignKey:SektorId" json:"sektor"`
	BankId               uint           `json:"bank_id"`
	Bank                 bank.Bank      `gorm:"foreignKey:BankCodeId" json:"bank"`
	InvestasiUser        []investasi.Investasi
}
