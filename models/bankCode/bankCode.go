package bankcode

import "os/user"

type BankCode struct {
	Id              uint        `gorm:"primaryKey" json:"id"`
	NamaBank        string      `json:"namaBank"`
	KodeBank        string      `json:"kodeBank"`
	ProyekMitraBank []user.User `gorm:"foreignKey:UserDetailsId" json:"proyekMitra" `
}
