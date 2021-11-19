package models

func (BankCode) TableName() string {
	return "bank_code"
}

type BankCode struct {
	Id         uint          `gorm:"primaryKey" json:"id"`
	NamaBank   string        `json:"namaBank"`
	KodeBank   string        `json:"kodeBank"`
	ProyekBank []ProyekMitra `json:"proyek_bank"`
}
