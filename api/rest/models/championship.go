package models

// This text will appear as description of your response body.
// swagger:response GetChampionshipResponse
type GetChampionshipResponseWrapper struct {
	// in: body
	Body GetChampionshipResponse
}

type GetChampionshipResponse struct {
	Championship ChampionshipDTO `json:"championship"`
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
