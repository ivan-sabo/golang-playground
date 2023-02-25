package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Championships []Championship

type Championship struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (c Championship) RegisterSeason(s Season) ChampionshipSeason {
	return ChampionshipSeason{
		championship: c,
		season:       s,
	}
}

func NewChampionship(name string) Championship {
	return Championship{
		ID:   uuid.NewString(),
		Name: name,
	}
}

type ChampionshipSeason struct {
	championship Championship
	season       Season
	clubs        Clubs
	rounds       Rounds
}

func (cs *ChampionshipSeason) AddClub(c Club) {
	cs.clubs = append(cs.clubs, c)
}

func (cs *ChampionshipSeason) RemoveClub(c Club) {
	for n, cl := range cs.clubs {
		if cl.ID == c.ID {
			cs.clubs[n] = cs.clubs[len(cs.clubs)-1]
			cs.clubs = cs.clubs[:len(cs.clubs)-1]
			break
		}
	}

	for i := range cs.rounds {
		cs.rounds[i].removeClub(c)
	}
}

func (cs *ChampionshipSeason) AddRound(r Round) {
	cs.rounds = append(cs.rounds, r)
}

func (cs *ChampionshipSeason) RemoveRound(r Round) {
	for n, ro := range cs.rounds {
		if ro.NO == r.NO {
			cs.rounds[n] = cs.rounds[len(cs.rounds)-1]
			cs.rounds = cs.rounds[:len(cs.rounds)-1]
			break
		}
	}
}

type Season struct {
	ID        string
	StartYear int
	EndYear   int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewSeason(s int, e int) Season {
	return Season{
		ID:        uuid.NewString(),
		StartYear: s,
		EndYear:   e,
	}
}

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

type Games []Game

type Game struct {
	ID        string
	Status    Status
	Home      Club
	Away      Club
	HomeScore int
	AwayScore int
	Start     time.Time
}

func (g *Game) Finish() {
	g.Status = Finished
}

func (g *Game) UpdateScore(home int, away int) {
	g.HomeScore = home
	g.AwayScore = away
}

func (g *Game) Schedule(start time.Time) {
	g.Start = start
	g.Status = Scheduled
}

func NewGame(home Club, away Club) Game {
	return Game{
		ID:     uuid.NewString(),
		Home:   home,
		Away:   away,
		Status: Created,
	}
}

type Rounds []Round

type Round struct {
	NO    int
	Games Games
}

func (r *Round) AddGame(g Game) {
	r.Games = append(r.Games, g)
}

func (r *Round) RemoveGame(g Game) {
	for n, ga := range r.Games {
		if ga.ID == g.ID {
			r.Games[n] = r.Games[len(r.Games)-1]
			r.Games = r.Games[:len(r.Games)-1]
			break
		}
	}
}

func (r *Round) removeClub(c Club) {
	games := Games{}
	for _, g := range r.Games {
		if g.Home.ID == c.ID || g.Away.ID == c.ID {
			continue
		}
		games = append(games, g)
	}
	r.Games = games
}

type Status string

var (
	Created    Status = "created"
	Scheduled  Status = "scheduled"
	InProgress Status = "in progress"
	Finished   Status = "finished"
)

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
