package models

import (
	"time"

	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
)

// List of championships reponse
// swagger:response GetChampionshipsResponse
type GetChampionshipsResponseWrapper struct {
	// in: body
	Body GetChampionshipsResponse
}

type GetChampionshipsResponse struct {
	Championships ChampionshipsDTO `json:"championships"`
}

func NewGetChampionshipsResponse(cs domain.Championships) GetChampionshipsResponse {
	return GetChampionshipsResponse{
		Championships: NewChampionshipsDTO(cs),
	}
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

// Single championship response
// swagger:response GetChampionshipResponse
type GetChampionshipResponseWrapper struct {
	// in: body
	Body GetChampionshipResponse
}

type GetChampionshipResponse struct {
	Championship ChampionshipDTO `json:"championship"`
}

func NewGetChampionshipResponse(c domain.Championship) GetChampionshipResponse {
	return GetChampionshipResponse{
		Championship: NewChampionshipDTO(c),
	}
}

// Single Championship request
// swagger:parameters CreateChampionship
type PostChampionshipRequestWrapper struct {
	// in: body
	Body PostChampionshipRequest
}

type PostChampionshipRequest struct {
	Name string `json:"name"`
}

func (r PostChampionshipRequest) ToEntity() domain.Championship {
	return domain.Championship{
		Name: r.Name,
	}
}

// Create Championship response
// swagger:response PostChampionshipResponse
type PostChampionshipResponseWrapper struct {
	// in: body
	Body PostChampionshipResponse
}

type PostChampionshipResponse struct {
	Championship ChampionshipDTO
}

func NewPostChampionshipResponse(c domain.Championship) PostChampionshipResponse {
	return PostChampionshipResponse{
		Championship: NewChampionshipDTO(c),
	}
}
