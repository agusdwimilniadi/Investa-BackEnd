package models

func (Sektor) TableName() string {
	return "sektor"
}

type Sektor struct {
	ID           uint          `gorm:"primaryKey" json:"id"`
	NamaSektor   string        `json:"nama_sektor"`
	ProyekProyek []ProyekMitra `json:"proyek_proyek"`
}
