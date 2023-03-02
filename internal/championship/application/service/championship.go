package service

import (
	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
	"github.com/ivan-sabo/golang-playground/internal/championship/infrastructure/database/mysql/repository"
	"gorm.io/gorm"
)

type ChampionshipService struct {
	championshipRepo       domain.ChampionshipRepo
	championshipSeasonRepo domain.ChampionshipSeasonRepo
	seasonRepo             domain.SeasonRepo
}

func NewChampionshipService(conn *gorm.DB) *ChampionshipService {
	return &ChampionshipService{
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

func (cs *ChampionshipService) GetChampionships(f domain.ChampionshipFilter) (domain.Championships, error) {
	c, err := cs.championshipRepo.GetChampionships(f)
	if err != nil {
		return domain.Championships{}, err
	}

	return c, nil
}

func (cs *ChampionshipService) GetChampionship(id string) (domain.Championship, error) {
	c, err := cs.championshipRepo.GetChampionship(id)
	if err != nil {
		return domain.Championship{}, err
	}

	return c, nil
}

func (cs *ChampionshipService) CreateChampionship(c domain.Championship) (domain.Championship, error) {
	c, err := cs.championshipRepo.CreateChampionship(c)
	if err != nil {
		return domain.Championship{}, err
	}

	return c, nil
}

func (cs *ChampionshipService) DeleteChampionship(id string) error {
	err := cs.championshipRepo.DeleteChampionship(id)
	if err != nil {
		return err
	}

	return nil
}

func (cs *ChampionshipService) GetChampionshipsSeasons(f domain.ChampionshipSeasonFilter) (domain.ChampionshipsSeasons, error) {
	c, err := cs.championshipSeasonRepo.GetChampionshipsSeasons(f)
	if err != nil {
		return domain.ChampionshipsSeasons{}, err
	}

	return c, nil
}
