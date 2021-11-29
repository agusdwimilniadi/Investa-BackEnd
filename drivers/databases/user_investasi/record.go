package user_investasi

import (
	"investaBackend/business/user_investasi"
	"investaBackend/drivers/databases/investasi"

	"gorm.io/gorm"
)

type UserInvestasi struct {
	gorm.Model
	Name          string
	Email         string
	Password      string
	InvestasiUser []investasi.Investasi
}

func (user UserInvestasi) ToDomain() user_investasi.Domain {
	return user_investasi.Domain{
		Id:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		Email:     user.Email,
		Name:      user.Name,
		Password:  user.Password,
	}

}

func FromDomain(domain user_investasi.Domain) UserInvestasi {
	return UserInvestasi{
		Model: gorm.Model{
			ID:        domain.Id,
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
			DeletedAt: domain.DeletedAt,
		},
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
	}
}
