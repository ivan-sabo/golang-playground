package models

import "github.com/ivan-sabo/golang-playground/internal/championship/domain"

// This text will appear as description of your response body.
// swagger:response GetClubsResponse
type GetClubsResponseWrapper struct {
	// in: body
	Body GetClubsResponse
}

type GetClubsResponse struct {
	Clubs ClubsDTO `json:"clubs"`
}

type ClubsDTO []ClubDTO

type ClubDTO struct {
	// Name of club
	//
	// required: true
	Name string `json:"name"`
}

func NewClubDTO(c domain.Club) ClubDTO {
	return ClubDTO{
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
