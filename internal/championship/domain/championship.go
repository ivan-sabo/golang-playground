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
}

type ClubRepo interface {
	GetClubs(ClubFilter) (Clubs, error)
	GetClub(string) (Club, error)
	CreateClub(Club) (Club, error)
	//UpdateClub(Club) (Club, error)
	//DeleteClub(int) error
}

type ClubFilter struct {
	Name string
}

var ErrClubNotFound = errors.New("club not found")
