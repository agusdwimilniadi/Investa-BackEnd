package investasi

import (
	"investaBackend/business/investasi"

	"gorm.io/gorm"
)

type Investasi struct {
	gorm.Model
	UserInvestasiID uint
	ProyekID        uint
	Nominal         int
	LinkBuktiTf     string
	// Proyek          proyek_mitra.ProyekMitra `json:"proyek"`
}
type TotalInvestasi struct {
	Total int
}

func (totalInvest *TotalInvestasi) ToDomainTotal() investasi.DomainTotalInvestasi {
	return investasi.DomainTotalInvestasi{
		TotalInvestasi: totalInvest.Total,
	}
}
func (invest *Investasi) ToDomain() investasi.DomainInvestasi {
	return investasi.DomainInvestasi{
		Model: gorm.Model{
			ID:        invest.ID,
			CreatedAt: invest.CreatedAt,
			UpdatedAt: invest.UpdatedAt,
			DeletedAt: invest.DeletedAt,
		},
		UserInvestasiID: invest.UserInvestasiID,
		ProyekID:        invest.ProyekID,
		Nominal:         invest.Nominal,
		LinkBuktiTf:     invest.LinkBuktiTf,
	}
}

func FromDomainTotal(total investasi.DomainTotalInvestasi) TotalInvestasi {
	return TotalInvestasi{
		Total: total.TotalInvestasi,
	}
}
func FromDomain(invest investasi.DomainInvestasi) Investasi {
	return Investasi{
		Model: gorm.Model{
			ID:        invest.ID,
			CreatedAt: invest.CreatedAt,
			UpdatedAt: invest.UpdatedAt,
			DeletedAt: invest.DeletedAt,
		},
		UserInvestasiID: invest.UserInvestasiID,
		ProyekID:        invest.ProyekID,
		Nominal:         invest.Nominal,
		LinkBuktiTf:     invest.LinkBuktiTf,
	}
}
