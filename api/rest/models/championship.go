package models

// This text will appear as description of your response body.
// swagger:response championshipResponse
type ChampionshipResponseWrapper struct {
	// in:body
	Body ChampionshipResponse
}

// ChampionshipResponse represents a body of Championship reponse
type ChampionshipResponse struct {
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
