package service

import (
	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
	"github.com/ivan-sabo/golang-playground/internal/championship/infrastructure/database/mysql/repository"
	"gorm.io/gorm"
)

type ChampionshipService struct {
	championshipRepo domain.ChampionshipRepo
}

func NewChampionshipService(conn *gorm.DB) ChampionshipService {
	return ChampionshipService{
		championshipRepo: repository.NewChampionshipMySQLRepo(conn),
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
