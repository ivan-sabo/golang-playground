package models

import (
	"time"

	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
)

// List of championships
// swagger:response GetChampionshipsResponse
type GetChampionshipsResponseWrapper struct {
	// in: body
	Body GetChampionshipsResponse
}

type GetChampionshipsResponse struct {
	Championships ChampionshipsDTO `json:"championships"`
}

type ChampionshipsDTO []ChampionshipDTO

type ChampionshipDTO struct {
	// required: true
	ID string `json:"id"`

	// required: true
	Name string `json:"name"`

	CreatedAt time.Time `json:"created_at"`

	UpdateAt time.Time `json:"updated_at"`

	DeletedAt time.Time `json:"deleted_at"`
}

func NewChampionshipDTO(c domain.Championship) ChampionshipDTO {
	return ChampionshipDTO{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdateAt:  c.UpdatedAt,
		DeletedAt: c.DeletedAt,
	}
}

func NewChampionshipsDTO(cs domain.Championships) ChampionshipsDTO {
	championships := make(ChampionshipsDTO, 0, len(cs))
	for _, c := range cs {
		championships = append(championships, NewChampionshipDTO(c))
	}

	return championships
}
