package service

import (
	"context"

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

func (s *SeasonService) GetSeasons(ctx context.Context, f domain.SeasonFilter) (domain.Seasons, error) {
	se, err := s.seasonRepo.GetSeasons(ctx, f)
	if err != nil {
		return nil, err
	}

	return se, nil
}

func (s *SeasonService) GetSeason(ctx context.Context, id string) (domain.Season, error) {
	se, err := s.seasonRepo.GetSeason(ctx, id)
	if err != nil {
		return domain.Season{}, err
	}

	return se, nil
}

func (s *SeasonService) CreateSeason(ctx context.Context, ds domain.Season) (domain.Season, error) {
	se, err := s.seasonRepo.CreateSeason(ctx, ds)
	if err != nil {
		return domain.Season{}, err
	}

	return se, nil
}

func (s *SeasonService) DeleteSeason(ctx context.Context, id string) error {
	err := s.seasonRepo.DeleteSeason(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
