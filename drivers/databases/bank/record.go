package bank

import (
	"investaBackend/business/bank"
	proyekMitraDomain "investaBackend/business/proyek_mitra"

	"investaBackend/drivers/databases/proyek_mitra"
)

func (Bank) TableName() string {
	return "bank_code"
}

type Bank struct {
	Id            uint `gorm:"primaryKey"`
	NamaBank      string
	KodeBank      string
	ProyekMitraId uint
	ProyekBank    []proyek_mitra.ProyekMitra
}

func (banks *Bank) ToDomain() bank.Domain {
	var proyekBank []proyekMitraDomain.Domain
	for _, val := range banks.ProyekBank {
		proyekBank = append(proyekBank, val.ToDomain())
	}
	return bank.Domain{
		Id:         banks.Id,
		NamaBank:   banks.NamaBank,
		KodeBank:   banks.KodeBank,
		ProyekBank: proyekBank,
	}
}

func FromDomain(domains bank.Domain) Bank {
	return Bank{
		Id:       domains.Id,
		NamaBank: domains.NamaBank,
		KodeBank: domains.KodeBank,
	}
}
