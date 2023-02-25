package domain

import (
	"time"

	"github.com/google/uuid"
)

type Clubs []Club

type Club struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewClub(n string) Club {
	return Club{
		ID:   uuid.NewString(),
		Name: n,
	}
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
