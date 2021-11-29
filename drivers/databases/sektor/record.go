package sektor

import "investaBackend/drivers/databases/proyek_mitra"

func (Sektor) TableName() string {
	return "sektor"
}

type Sektor struct {
	ID           int `gorm:"primaryKey"`
	NamaSektor   string
	ProyekSektor []proyek_mitra.ProyekMitra
}
