package bankcode

type BankCode struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	NamaBank string `json:"namaBank"`
	KodeBank string `json:"kodeBank"`
	// ProyekMitraBank  []sektor.
}
