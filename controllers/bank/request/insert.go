package request

import "investaBackend/business/bank"

type InsertBank struct {
	NamaBank string `json:"nama_bank"`
	KodeBank string `json:"kode_bank"`
}

func (banks *InsertBank) ToDomain() bank.Domain {
	return bank.Domain{
		NamaBank: banks.NamaBank,
		KodeBank: banks.KodeBank,
	}
}
