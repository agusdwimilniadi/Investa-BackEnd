package bank

import "investaBackend/business/bank"

func (Bank) TableName() string {
	return "bank_code"
}

type Bank struct {
	Id       uint `gorm:"primaryKey"`
	NamaBank string
	KodeBank string
	// ProyekBank []ProyekMitra

}

func (banks *Bank) ToDomain() bank.Domain {
	return bank.Domain{
		Id:       banks.Id,
		NamaBank: banks.NamaBank,
		KodeBank: banks.KodeBank,
	}
}

func FromDomain(domains bank.Domain) Bank {
	return Bank{
		Id:       domains.Id,
		NamaBank: domains.NamaBank,
		KodeBank: domains.KodeBank,
	}
}
