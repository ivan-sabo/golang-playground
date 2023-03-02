package models

import "github.com/ivan-sabo/golang-playground/internal/championship/domain"

// Create ChampionshipSeason response
// swagger:response RegisterSeasonResponse
type RegisterSeasonResponseWrapper struct {
	// in: body
	Body RegisterSeasonResponse
}

type RegisterSeasonResponse struct {
	ChampionshipSeason ChampionshipSeasonDTO
}

func NewRegisterSeasonResponse(cs domain.ChampionshipSeason) ChampionshipSeasonDTO {
	return NewChampionshipSeason(cs)
}

type ChampionshipsSeasonsDTO []ChampionshipSeasonDTO

type ChampionshipSeasonDTO struct {
	Championship ChampionshipDTO
	Season       SeasonDTO
}

func NewChampionshipSeason(dcs domain.ChampionshipSeason) ChampionshipSeasonDTO {
	return ChampionshipSeasonDTO{
		Championship: NewChampionshipDTO(dcs.Championship),
		Season:       NewSeasonDTO(dcs.Season),
	}
}

// Get ChampionshipsSeasons response
// swagger:response GetChampionshipsSeasonsResponse
type GetChampionshipsSeasonsResponseWrapper struct {
	// in: body
	Body GetChampionshipsSeasonsResponse
}

type GetChampionshipsSeasonsResponse struct {
	ChampionshipsSeasons ChampionshipsSeasonsDTO
}

func NewChampionshipsSeasons(dcsss domain.ChampionshipsSeasons) ChampionshipsSeasonsDTO {
	csss := make(ChampionshipsSeasonsDTO, 0, len(dcsss))

	for _, dcs := range dcsss {
		csss = append(csss, NewChampionshipSeason(dcs))
	}

	return csss
}
