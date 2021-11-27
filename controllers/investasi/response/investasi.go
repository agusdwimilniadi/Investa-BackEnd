package response

import (
	"investaBackend/business/investasi"

	"gorm.io/gorm"
)

type InvestasiResponse struct {
	gorm.Model
	UserInvestasiID uint   `json:"user_investasi_id"`
	ProyekID        uint   `json:"proyek_id"`
	Nominal         int    `json:"nominal"`
	LinkBuktiTf     string `json:"link_bukti_tf"`
}

func FromDomain(domain investasi.DomainInvestasi) InvestasiResponse {
	return InvestasiResponse{
		Model: gorm.Model{
			ID:        domain.ID,
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
			DeletedAt: domain.DeletedAt,
		},
		UserInvestasiID: domain.UserInvestasiID,
		ProyekID:        domain.ProyekID,
		Nominal:         domain.Nominal,
		LinkBuktiTf:     domain.LinkBuktiTf,
	}
}
