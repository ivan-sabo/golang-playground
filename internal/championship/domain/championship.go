package domain

import (
	"errors"
	"time"
)

type Championships []Championship

type Championship struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Seasons []Season

type Season struct {
	ID        string
	StartYear int
	EndYear   int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Clubs []Club

type Club struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type ChampionshipRepo interface {
	GetChampionships() (Championships, error)
	GetChampionship(string) (Championship, error)
	CreateChampionship(Championship) (Championship, error)
	UpdateChampionship(string, Championship) (Championship, error)
	DeleteChampionship(string) error
}

type ClubRepo interface {
	GetClubs(ClubFilter) (Clubs, error)
	GetClub(string) (Club, error)
	CreateClub(Club) (Club, error)
	UpdateClub(string, Club) (Club, error)
	DeleteClub(string) error
}

type ClubFilter struct {
	Name string
}

var (
	ErrClubNotFound         = errors.New("club not found")
	ErrChampionshipNotFound = errors.New("championship not found")
)
