package service

import (
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

func (s *ClubService) GetClubs(f domain.ClubFilter) (domain.Clubs, error) {
	c, err := s.clubRepo.GetClubs(f)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *ClubService) GetClub(id string) (domain.Club, error) {
	c, err := s.clubRepo.GetClub(id)
	if err != nil {
		return domain.Club{}, err
	}

	return c, nil
}

func (s *ClubService) CreateClub(c domain.Club) (domain.Club, error) {
	c, err := s.clubRepo.CreateClub(c)
	if err != nil {
		return domain.Club{}, err
	}

	return c, nil
}

func (s *ClubService) UpdateClub(id string, c domain.Club) (domain.Club, error) {
	c, err := s.clubRepo.UpdateClub(id, c)
	if err != nil {
		return domain.Club{}, err
	}

	return c, nil
}

func (s *ClubService) DeleteClub(id string) error {
	err := s.clubRepo.DeleteClub(id)
	if err != nil {
		return err
	}

	return nil
}
