package domain

import "time"

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
	GetClubs() (Clubs, error)
	GetClub(ClubFilter) (Club, error)
	CreateClub(Club) (Club, error)
	UpdateClub(Club) (Club, error)
	DeleteClub(int) error
}

type ClubFilter struct {
	Name string
	ID   string
}
