package models

import "github.com/ivan-sabo/golang-playground/internal/championship/domain"

// List of Seasons response.
// swagger:response GetSeasonsResponse
type GetSeasonsResponseWrapper struct {
	// in: body
	Body GetSeasonsResponse
}

type GetSeasonsResponse struct {
	Seasons SeasonsDTO `json:"seasons"`
}

func NewGetSeasonsResponse(c domain.Clubs) GetClubsResponse {
	return GetClubsResponse{
		Clubs: NewClubsDTO(c),
	}
}

type SeasonsDTO []SeasonDTO

type SeasonDTO struct {
	ID        string
	StartYear int
	EndYear   int
}
