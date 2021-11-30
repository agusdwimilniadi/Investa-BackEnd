package request

import (
	"investaBackend/business/investasi"

	"gorm.io/gorm"
)

type InsertInvestasiReq struct {
	gorm.Model
	UserInvestasiID uint   `json:"user_investasi_id"`
	ProyekMitraID   uint   `json:"proyek_id"`
	Nominal         int    `json:"nominal"`
	LinkBuktiTf     string `json:"link_bukti_tf"`
}

func (insert *InsertInvestasiReq) ToDomainReq() investasi.DomainInvestasi {
	return investasi.DomainInvestasi{
		Model: gorm.Model{
			ID:        insert.ID,
			CreatedAt: insert.CreatedAt,
			UpdatedAt: insert.UpdatedAt,
			DeletedAt: insert.DeletedAt,
		},
		UserInvestasiID: insert.UserInvestasiID,
		ProyekMitraID:   insert.ProyekMitraID,
		Nominal:         insert.Nominal,
		LinkBuktiTf:     insert.LinkBuktiTf,
	}
}
