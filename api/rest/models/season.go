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

func NewGetSeasonsResponse(c domain.Seasons) GetSeasonsResponse {
	return GetSeasonsResponse{
		Seasons: NewSeasonsDTO(c),
	}
}

type SeasonsDTO []SeasonDTO

func NewSeasonsDTO(ss domain.Seasons) SeasonsDTO {
	seasons := make(SeasonsDTO, 0, len(ss))
	for _, s := range ss {
		seasons = append(seasons, NewSeasonDTO(s))
	}
	return seasons
}

type SeasonDTO struct {
	ID        string
	StartYear int
	EndYear   int
}

func NewSeasonDTO(s domain.Season) SeasonDTO {
	return SeasonDTO{
		ID:        s.ID.String(),
		StartYear: s.StartYear,
		EndYear:   s.EndYear,
	}
}

// Create Season request
// swagger:parameters CreateSeason
type PostCSeasonRequestWrapper struct {
	// in: body
	Body PostSeasonRequest
}

type PostSeasonRequest struct {
	StartYear int `json:"startYear"`
	EndYear   int `json:"endYear"`
}

func (r PostSeasonRequest) ToEntity() (domain.Season, error) {
	ds, err := domain.NewSeason("", r.StartYear, r.EndYear)
	if err != nil {
		return domain.Season{}, err
	}

	return ds, nil
}

// Create Season response
// swagger:response PostSeasonResponse
type PostSeasonResponseWrapper struct {
	// in: body
	Body PostSeasonResponse
}

type PostSeasonResponse struct {
	Season SeasonDTO
}

func NewPostSeasonResponse(s domain.Season) PostSeasonResponse {
	return PostSeasonResponse{
		Season: NewSeasonDTO(s),
	}
}
