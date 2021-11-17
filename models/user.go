package models

import (
	"time"

	"gorm.io/gorm"
)

func (User) TableName() string {
	return "user"
}

type User struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Password   string         `json:"password"`
	ProyekUser []ProyekMitra  `json:"proyek_user"`
	// ProyekMitra []user.User    `gorm:"foreignKey:UserDetailsId" json:"proyekMitra"`
}
