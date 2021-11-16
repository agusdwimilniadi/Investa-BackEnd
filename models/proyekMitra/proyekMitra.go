package proyekmitra

import (
	"time"

	"gorm.io/gorm"
)

type ProyekMitra struct {
	Id                   uint           `gorm:"primaryKey" json:"id"`
	CreatedAt            time.Time      `json:"createdAt"`
	UpdatedAt            time.Time      `json:"updatedAt"`
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	NamaKelompokTani     string         `json:"namaKelompokTani"`
	NamaProyek           string         `json:"namaProyek"`
	PengalamanBertani    int            `json:"pengalamanBertani"`
	PersentaseKeuntungan int            `json:"persentaseKeuntungan"`
	NominalPengajuan     int            `json:"nominalPengajuan"`
	PeriodePanen         int            `json:"periodePanen"`
	AtasNamaRekening     string         `json:"atasNamaRekening"`
	DeskripsiProyek      string         `json:"deskripsiProyek"`
	DeskripsiUmum        string         `json:"deskripsiUmum"`
	StatusMitra          bool           `json:"statusMitra"`
	LinkDokumen          string         `json:"linkDokumen"`
	LinkFoto             string         `json:"linkFoto"`
	UserDetailsId        uint           `json:"userDetailsId"`
	SektorPengajuanId    uint           `json:"sektorPengajuanId"`
	BankId               uint           `json:"bankId"`
}
