package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/ivan-sabo/golang-playground/internal"
)

type Status string

type Championships []Championship

type Championship struct {
	ID        uuid.UUID
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

func NewChampionship(id string, name string) (Championship, error) {
	if id == "" {
		id = uuid.NewString()
	}

	cuuid, err := uuid.Parse(id)
	if err != nil {
		return Championship{}, internal.ErrInvalidUUID
	}

	return Championship{
		ID:   cuuid,
		Name: name,
	}, nil
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

type Seasons []Season

type Season struct {
	ID        uuid.UUID
	StartYear int
	EndYear   int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewSeason(id string, start int, end int) (Season, error) {
	if id == "" {
		id = uuid.NewString()
	}

	suuid, err := uuid.Parse(id)
	if err != nil {
		return Season{}, internal.ErrInvalidUUID
	}

	if (end - start) != 1 {
		return Season{}, ErrSeasonDurationInvalid
	}

	return Season{
		ID:        suuid,
		StartYear: start,
		EndYear:   end,
	}, nil
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

type SeasonFilter struct {
	StartYear int
	EndYear   int
}

type ChampionshipRepo interface {
	GetChampionships() (Championships, error)
	GetChampionship(string) (Championship, error)
	CreateChampionship(Championship) (Championship, error)
	UpdateChampionship(string, Championship) (Championship, error)
	DeleteChampionship(string) error
}

type SeasonRepo interface {
	GetSeasons(SeasonFilter) (Seasons, error)
	//GetSeason(string) (Season, error)
	CreateSeason(Season) (Season, error)
	//UpdateSeason(string, Season) (Season, error)
	//DeleteSeason(string) error
}

var (
	ErrClubNotFound          = errors.New("club not found")
	ErrChampionshipNotFound  = errors.New("championship not found")
	ErrSeasonDurationInvalid = errors.New("invalid season start and/or end values")
)
