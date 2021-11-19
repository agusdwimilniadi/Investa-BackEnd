package models

import (
	"time"

	"gorm.io/gorm"
)

// func (User) TableName() string {
// 	return "user"
// }

type UserProyek struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Password   string         `json:"password"`
	ProyekUser []ProyekMitra  `json:"proyek_user"`
}
