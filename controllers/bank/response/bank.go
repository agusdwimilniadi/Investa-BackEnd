package response

import "investaBackend/business/bank"

type InsertBankResponse struct {
	Id       uint   `json:"id"`
	NamaBank string `json:"nama_bank"`
	KodeBank string `json:"kode_bank"`
}

func FromDomain(domain bank.Domain) InsertBankResponse {
	return InsertBankResponse{
		Id:       domain.Id,
		NamaBank: domain.NamaBank,
		KodeBank: domain.KodeBank,
	}
}
