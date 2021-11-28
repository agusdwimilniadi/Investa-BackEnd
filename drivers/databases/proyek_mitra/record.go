package proyek_mitra

import (
	"investaBackend/business/proyek_mitra"

	"gorm.io/gorm"
)

func (ProyekMitra) TableName() string {
	return "proyek_mitra"
}

type ProyekMitra struct {
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

	// Sektor               sektor.Sektor  `gorm:"foreignKey:SektorId" json:"sektor"`
	// Bank                 bank.Bank      `gorm:"foreignKey:BankCodeId" json:"bank"`
	// InvestasiUser        []investasi.Investasi
}

func (getData *ProyekMitra) ToDomain() proyek_mitra.Domain {
	return proyek_mitra.Domain{
		Model: gorm.Model{
			ID:        getData.ID,
			CreatedAt: getData.CreatedAt,
			UpdatedAt: getData.CreatedAt,
			DeletedAt: getData.DeletedAt,
		},
		NamaKelompokTani:     getData.NamaKelompokTani,
		NamaProyek:           getData.NamaProyek,
		PengalamanBertani:    getData.PengalamanBertani,
		PersentaseKeuntungan: getData.PersentaseKeuntungan,
		NominalPengajuan:     getData.NominalPengajuan,
		PeriodePanen:         getData.PeriodePanen,
		AtasNamaRekening:     getData.AtasNamaRekening,
		DeskripsiProyek:      getData.AtasNamaRekening,
		DeskripsiUmum:        getData.DeskripsiUmum,
		StatusMitra:          getData.StatusMitra,
		LinkDokumen:          getData.LinkDokumen,
		LinkFoto:             getData.LinkFoto,
		UserID:               getData.UserID,
		SektorId:             getData.SektorId,
		BankId:               getData.BankId,
	}
}
func FromDomain(getData proyek_mitra.Domain) ProyekMitra {
	return ProyekMitra{
		Model: gorm.Model{
			ID:        getData.ID,
			CreatedAt: getData.CreatedAt,
			UpdatedAt: getData.CreatedAt,
			DeletedAt: getData.DeletedAt,
		},
		NamaKelompokTani:     getData.NamaKelompokTani,
		NamaProyek:           getData.NamaProyek,
		PengalamanBertani:    getData.PengalamanBertani,
		PersentaseKeuntungan: getData.PersentaseKeuntungan,
		NominalPengajuan:     getData.NominalPengajuan,
		PeriodePanen:         getData.PeriodePanen,
		AtasNamaRekening:     getData.AtasNamaRekening,
		DeskripsiProyek:      getData.AtasNamaRekening,
		DeskripsiUmum:        getData.DeskripsiUmum,
		StatusMitra:          getData.StatusMitra,
		LinkDokumen:          getData.LinkDokumen,
		LinkFoto:             getData.LinkFoto,
		UserID:               getData.UserID,
		SektorId:             getData.SektorId,
		BankId:               getData.BankId,
	}
}

func ToDomains(u []ProyekMitra) []proyek_mitra.Domain {
	var Domains []proyek_mitra.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
