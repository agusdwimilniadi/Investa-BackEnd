package request

import "investaBackend/business/sektor"

type InsertSektor struct {
	NamaSektor string `json:"nama_sektor"`
}

func (sektors *InsertSektor) ToDomain() sektor.Domain {
	return sektor.Domain{
		NamaSektor: sektors.NamaSektor,
	}
}
