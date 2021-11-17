package models

type Sektor struct {
	ID           uint          `gorm:"primaryKey" json:"id"`
	NamaSektor   string        `json:"nama_sektor"`
	ProyekProyek []ProyekMitra `json:"proyek_proyek"`
}
