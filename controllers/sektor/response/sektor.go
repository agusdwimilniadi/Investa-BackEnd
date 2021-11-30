package response

import "investaBackend/business/sektor"

type InsertSektorResponse struct {
	Id         uint   `json:"id"`
	NamaSektor string `json:"nama_sektor`
}

func FromDomain(domain sektor.Domain) InsertSektorResponse {
	return InsertSektorResponse{
		Id:         domain.Id,
		NamaSektor: domain.NamaSektor,
	}

}
