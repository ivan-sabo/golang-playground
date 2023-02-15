package models

// This text will appear as description of your response body.
// swagger:response ChampionshipResponse
type ChampionshipResponseWrapper struct {
	// in: body
	Body Championship
}

// Championship represents a body of Championship reponse
// swagger:model
type Championship struct {
	// Collection of clubs
	// required: true
	Clubs Clubs `json:"clubs"`
}

// Clubs is a collection of Club type

type Clubs []Club

// Club represents a football club
type Club struct {
	// Name of club
	//
	// required: true
	Name string `json:"name"`
}
