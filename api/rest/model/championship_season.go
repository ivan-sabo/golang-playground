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
