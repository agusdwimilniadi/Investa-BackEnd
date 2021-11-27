package models

import (
	"time"
)

// func (DetailInvestasi) TableName() string {
// 	return "detail_investasi"
// }

type Investasi struct {
	ID              uint        `gorm:"primaryKey" json:"id"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	UserInvestasiID uint        `json:"user_investasi_id"`
	ProyekID        uint        `json:"proyek_id"`
	Proyek          ProyekMitra `json:"proyek"`
	Nominal         int         `json:"nominal"`
	LinkBuktiTf     string      `json:"link_bukti_tf"`
}
