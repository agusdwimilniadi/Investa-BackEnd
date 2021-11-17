package models

import (
	"time"
)

type DetailInvestasi struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	UserID      uint        `json:"user_id"`
	User        User        `json:"user"`
	ProyekID    uint        `json:"proyek_id"`
	Proyek      ProyekMitra `json:"proyek"`
	Nominal     int         `json:"nominal"`
	LinkBuktiTf string      `json:"link_bukti_tf"`
}
