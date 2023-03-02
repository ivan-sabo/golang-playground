package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/ivan-sabo/golang-playground/internal"
)

type Clubs []Club

type Club struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewClub(id string, name string) (Club, error) {
	if id == "" {
		id = uuid.NewString()
	}

	cuuid, err := uuid.Parse(id)
	if err != nil {
		return Club{}, internal.ErrInvalidUUID
	}

	return Club{
		ID:   cuuid,
		Name: name,
	}, nil
}

type ClubFilter struct {
	Name string
}

type ClubRepo interface {
	GetClubs(ClubFilter) (Clubs, error)
	GetClub(string) (Club, error)
	CreateClub(Club) (Club, error)
	UpdateClub(string, Club) (Club, error)
	DeleteClub(string) error
}

type ClubService interface {
	GetClubs(ClubFilter) (Clubs, error)
	GetClub(string) (Club, error)
	CreateClub(Club) (Club, error)
	UpdateClub(string, Club) (Club, error)
	DeleteClub(string) error
}
