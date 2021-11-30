package sektor

import (
	"investaBackend/business/sektor"
	"investaBackend/drivers/databases/proyek_mitra"
)

func (Sektor) TableName() string {
	return "sektor"
}

type Sektor struct {
	ID           int `gorm:"primaryKey"`
	NamaSektor   string
	ProyekSektor []proyek_mitra.ProyekMitra
}

func (sektors *Sektor) ToDomain() sektor.Domain {
	return sektor.Domain{
		Id:         uint(sektors.ID),
		NamaSektor: sektors.NamaSektor,
	}
}

func FromDomain(domains sektor.Domain) Sektor {
	return Sektor{
		ID:         int(domains.Id),
		NamaSektor: domains.NamaSektor,
		// ProyekSektor: []proyek_mitra.ProyekMitra{},
	}
}
