package detailinvestasi

import "time"

type DetailInvestasi struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UserId      string    `json:"userId"`
	Nominal     int       `json:"nominal"`
	LinkBuktiTf string    `json:"linkBuktiTf"`
}
