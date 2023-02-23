package models

import "github.com/ivan-sabo/golang-playground/internal/championship/domain"

// List of Clubs response.
// swagger:response GetClubsResponse
type GetClubsResponseWrapper struct {
	// in: body
	Body GetClubsResponse
}

type GetClubsResponse struct {
	Clubs ClubsDTO `json:"clubs"`
}

// Single Club response
// swagger:response GetClubResponse
type GetClubResponseWrapper struct {
	// in: body
	Body GetClubResponse
}

type GetClubResponse struct {
	Club ClubDTO `json:"club"`
}

type ClubsDTO []ClubDTO

type ClubDTO struct {
	ID string `json:"id"`

	// required: true
	Name string `json:"name"`
}

func NewClubDTO(c domain.Club) ClubDTO {
	return ClubDTO{
		ID:   c.ID,
		Name: c.Name,
	}
}

func NewClubsDTO(cs domain.Clubs) ClubsDTO {
	clubs := make(ClubsDTO, 0, len(cs))
	for _, c := range cs {
		clubs = append(clubs, NewClubDTO(c))
	}

	return clubs
}

// Single Club request
// swagger:parameters CreateClub
type PostClubRequestWrapper struct {
	// in: body
	Body PostClubRequest
}

type PostClubRequest struct {
	Name string `json:"name"`
}

func (r PostClubRequest) ToEntity() domain.Club {
	return domain.Club{
		Name: r.Name,
	}
}

// Create Club response
// swagger:response PostClubResponse
type PostClubResponseWrapper struct {
	// in: body
	Body PostClubResponse
}

type PostClubResponse struct {
	Name string `json:"name"`
}
