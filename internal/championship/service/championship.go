package service

import (
	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
	"github.com/ivan-sabo/golang-playground/internal/championship/infrastructure/database/mysql/repository"
	"gorm.io/gorm"
)

type ChampionshipService struct {
	championshipRepo       domain.ChampionshipRepo
	seasonRepo             domain.SeasonRepo
	championshipSeasonRepo domain.ChampionshipSeasonRepo
}

func NewChampionshipService(conn *gorm.DB) ChampionshipService {
	return ChampionshipService{
		championshipRepo:       repository.NewChampionshipMySQLRepo(conn),
		seasonRepo:             repository.NewSeasonMySQLRepo(conn),
		championshipSeasonRepo: repository.NewChampionshipSeasonMySQLRepo(conn),
	}
}

func (cs *ChampionshipService) UpdateChampionship(id string, c domain.Championship) (domain.Championship, error) {
	_, err := cs.championshipRepo.GetChampionship(id)
	if err != nil {
		return domain.Championship{}, err
	}

	dc, err := cs.championshipRepo.UpdateChampionship(id, c)
	if err != nil {
		return domain.Championship{}, err
	}

	return dc, nil
}

func (cs *ChampionshipService) RegisterSeason(championshipID string, seasonID string) (domain.ChampionshipSeason, error) {
	championship, err := cs.championshipRepo.GetChampionship(championshipID)
	if err != nil {
		return domain.ChampionshipSeason{}, err
	}

	season, err := cs.seasonRepo.GetSeason(seasonID)
	if err != nil {
		return domain.ChampionshipSeason{}, err
	}

	ncs := championship.RegisterSeason(season)
	dcs, err := cs.championshipSeasonRepo.RegisterSeason(ncs)
	if err != nil {
		return domain.ChampionshipSeason{}, err
	}

	return dcs, nil
}
