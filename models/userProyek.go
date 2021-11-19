package models

import (
	"gorm.io/gorm"
)

func (UserProyek) TableName() string {
	return "user_proyek"
}

type UserProyek struct {
	gorm.Model
	Name       string        `json:"name"`
	Email      string        `json:"email"`
	Password   string        `json:"password"`
	ProyekUser []ProyekMitra `json:"proyek_user"`
}
