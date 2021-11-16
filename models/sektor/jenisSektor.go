package sektor

type Sektor struct {
	Id         uint   `gorm:"primaryKey" json:"id"`
	NamaSektor string `json:"NamaSektor"`
}
