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

func NewGetClubsResponse(c domain.Clubs) GetClubsResponse {
	return GetClubsResponse{
		Clubs: NewClubsDTO(c),
	}
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

func NewGetClubResponse(c domain.Club) GetClubResponse {
	return GetClubResponse{
		Club: NewClubDTO(c),
	}
}

type ClubsDTO []ClubDTO

func NewClubsDTO(cs domain.Clubs) ClubsDTO {
	clubs := make(ClubsDTO, 0, len(cs))
	for _, c := range cs {
		clubs = append(clubs, NewClubDTO(c))
	}

	return clubs
}

type ClubDTO struct {
	ID string `json:"id"`

	// required: true
	Name string `json:"name"`
}

func NewClubDTO(c domain.Club) ClubDTO {
	return ClubDTO{
		ID:   c.ID.String(),
		Name: c.Name,
	}
}

// Create Club request
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
	Club ClubDTO
}

func NewPostClubResponse(c domain.Club) PostClubResponse {
	return PostClubResponse{
		Club: NewClubDTO(c),
	}
}

// Update Club request
// swagger:parameters UpdateClub
type PutClubRequestWrapper struct {
	// in: body
	Body PutClubRequest
}

type PutClubRequest struct {
	Name string `json:"name"`
}

func (r PutClubRequest) ToEntity() domain.Club {
	return domain.Club{
		Name: r.Name,
	}
}

// Update Club response
// swagger:response PostClubResponse
type PutClubResponseWrapper struct {
	// in: body
	Body PutClubResponse
}

type PutClubResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewPutClubResponse(c domain.Club) PutClubResponse {
	return PutClubResponse{
		ID:   c.ID.String(),
		Name: c.Name,
	}
}
