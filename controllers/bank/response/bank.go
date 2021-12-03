package response

import (
	"investaBackend/business/bank"
	"investaBackend/business/proyek_mitra"
)

type InsertBankResponse struct {
	Id         uint   `json:"id"`
	NamaBank   string `json:"nama_bank"`
	KodeBank   string `json:"kode_bank"`
	ProyekBank []ProyekMitra
}

func FromDomain(domain bank.Domain) InsertBankResponse {
	var proyekBank []ProyekMitra
	for _, val := range domain.ProyekBank {
		proyekBank = append(proyekBank, proyekFromDomain(val))
	}
	return InsertBankResponse{
		Id:         domain.Id,
		NamaBank:   domain.NamaBank,
		KodeBank:   domain.KodeBank,
		ProyekBank: proyekBank,
	}
}

// func FromDomainList(domain []bank.Domain) []InsertBankResponse {
// 	var list []InsertBankResponse
// 	for _, k := range domain {
// 		tempResp := InsertBankResponse{
// 			Id:       k.Id,
// 			NamaBank: k.NamaBank,
// 			KodeBank: k.KodeBank,
// 		}
// 		list = append(list, tempResp)
// 	}
// 	return list
// }

type ProyekMitra struct {
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

func proyekFromDomain(u proyek_mitra.Domain) ProyekMitra {
	return ProyekMitra{
		NamaKelompokTani:     u.NamaKelompokTani,
		NamaProyek:           u.NamaProyek,
		PengalamanBertani:    u.PengalamanBertani,
		PersentaseKeuntungan: u.PersentaseKeuntungan,
		NominalPengajuan:     u.NominalPengajuan,
		PeriodePanen:         u.PeriodePanen,
		AtasNamaRekening:     u.AtasNamaRekening,
		DeskripsiProyek:      u.DeskripsiProyek,
		DeskripsiUmum:        u.DeskripsiUmum,
		StatusMitra:          u.StatusMitra,
		LinkDokumen:          u.LinkDokumen,
		LinkFoto:             u.LinkFoto,
		UserID:               u.ID,
		SektorId:             u.SektorId,
		BankId:               u.BankId,
		ProyekID:             u.ProyekID,
	}
}
