package service

import (
	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
	"github.com/ivan-sabo/golang-playground/internal/championship/infrastructure/database/mysql/repository"
	"gorm.io/gorm"
)

type SeasonService struct {
	seasonRepo domain.SeasonRepo
}

func NewSeasonService(conn *gorm.DB) *SeasonService {
	return &SeasonService{
		seasonRepo: repository.NewSeasonMySQLRepo(conn),
	}
}

func (s *SeasonService) GetSeasons(f domain.SeasonFilter) (domain.Seasons, error) {
	se, err := s.seasonRepo.GetSeasons(f)
	if err != nil {
		return nil, err
	}

	return se, nil
}

func (s *SeasonService) GetSeason(id string) (domain.Season, error) {
	se, err := s.seasonRepo.GetSeason(id)
	if err != nil {
		return domain.Season{}, err
	}

	return se, nil
}

func (s *SeasonService) CreateSeason(ds domain.Season) (domain.Season, error) {
	se, err := s.seasonRepo.CreateSeason(ds)
	if err != nil {
		return domain.Season{}, err
	}

	return se, nil
}

func (s *SeasonService) DeleteSeason(id string) error {
	err := s.seasonRepo.DeleteSeason(id)
	if err != nil {
		return err
	}

	return nil
}
