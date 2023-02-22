package models

import "github.com/ivan-sabo/golang-playground/internal/championship/domain"

// This text will appear as description of your response body.
// swagger:response GetChampionshipResponse
type GetChampionshipResponseWrapper struct {
	// in: body
	Body GetChampionshipResponse
}

type GetChampionshipResponse struct {
	Championship ChampionshipDTO `json:"championship"`
}

func NewChampionshipDTO(domain.Championship) ChampionshipDTO {
	return ChampionshipDTO{}
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

// Championship represents a body of Championship reponse
type ChampionshipDTO struct {
	// Collection of clubs
	// required: true
	Clubs ClubsDTO `json:"clubs"`

	// Name of the championship
	Name string `json:"name"`
}

// Clubs is a collection of Club type
type ClubsDTO []ClubDTO

// Club represents a football club
type ClubDTO struct {
	// Name of club
	//
	// required: true
	Name string `json:"name"`
}
