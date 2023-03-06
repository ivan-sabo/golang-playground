package service

import (
	"context"

	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
	"github.com/ivan-sabo/golang-playground/internal/championship/infrastructure/database/mysql/repository"
	"gorm.io/gorm"
)

type ClubService struct {
	clubRepo domain.ClubRepo
}

func NewClubService(conn *gorm.DB) *ClubService {
	return &ClubService{
		clubRepo: repository.NewClubMySQLRepo(conn),
	}
}

func (s *ClubService) GetClubs(ctx context.Context, f domain.ClubFilter) (domain.Clubs, error) {
	c, err := s.clubRepo.GetClubs(ctx, f)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *ClubService) GetClub(ctx context.Context, id string) (domain.Club, error) {
	c, err := s.clubRepo.GetClub(ctx, id)
	if err != nil {
		return domain.Club{}, err
	}

	return c, nil
}

func (s *ClubService) CreateClub(ctx context.Context, c domain.Club) (domain.Club, error) {
	c, err := s.clubRepo.CreateClub(ctx, c)
	if err != nil {
		return domain.Club{}, err
	}

	return c, nil
}

func (s *ClubService) UpdateClub(ctx context.Context, id string, c domain.Club) (domain.Club, error) {
	c, err := s.clubRepo.UpdateClub(ctx, id, c)
	if err != nil {
		return domain.Club{}, err
	}

	return c, nil
}

func (s *ClubService) DeleteClub(ctx context.Context, id string) error {
	err := s.clubRepo.DeleteClub(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
