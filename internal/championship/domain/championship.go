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
		Championship: c,
		Season:       s,
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

type ChampionshipsSeasons []ChampionshipSeason

type ChampionshipSeason struct {
	Championship Championship
	Season       Season
	Clubs        Clubs
	Rounds       Rounds
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

func (cs *ChampionshipSeason) AddClub(c Club) {
	cs.Clubs = append(cs.Clubs, c)
}

func (cs *ChampionshipSeason) RemoveClub(c Club) {
	for n, cl := range cs.Clubs {
		if cl.ID == c.ID {
			cs.Clubs[n] = cs.Clubs[len(cs.Clubs)-1]
			cs.Clubs = cs.Clubs[:len(cs.Clubs)-1]
			break
		}
	}

	for i := range cs.Rounds {
		cs.Rounds[i].removeClub(c)
	}
}

func (cs *ChampionshipSeason) AddRound(r Round) {
	cs.Rounds = append(cs.Rounds, r)
}

func (cs *ChampionshipSeason) RemoveRound(r Round) {
	for n, ro := range cs.Rounds {
		if ro.NO == r.NO {
			cs.Rounds[n] = cs.Rounds[len(cs.Rounds)-1]
			cs.Rounds = cs.Rounds[:len(cs.Rounds)-1]
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

type ChampionshipFilter struct {
	Name string
}

type SeasonFilter struct {
	StartYear int
	EndYear   int
}

type ChampionshipSeasonFilter struct {
	ChampionshipID string
	SeasonID       string
	IncludeDeleted bool
}

type ChampionshipRepo interface {
	GetChampionships(ChampionshipFilter) (Championships, error)
	GetChampionship(string) (Championship, error)
	CreateChampionship(Championship) (Championship, error)
	UpdateChampionship(string, Championship) (Championship, error)
	DeleteChampionship(string) error
}

type SeasonRepo interface {
	GetSeasons(SeasonFilter) (Seasons, error)
	GetSeason(string) (Season, error)
	CreateSeason(Season) (Season, error)
	DeleteSeason(string) error
}

type ChampionshipSeasonRepo interface {
	RegisterSeason(ChampionshipSeason) (ChampionshipSeason, error)
	GetChampionshipsSeasons(ChampionshipSeasonFilter) (ChampionshipsSeasons, error)
}

type ChampionshipService interface {
	GetChampionships(ChampionshipFilter) (Championships, error)
	GetChampionship(string) (Championship, error)
	CreateChampionship(Championship) (Championship, error)
	UpdateChampionship(string, Championship) (Championship, error)
	DeleteChampionship(string) error
	RegisterSeason(string, string) (ChampionshipSeason, error)
	GetChampionshipsSeasons(ChampionshipSeasonFilter) (ChampionshipsSeasons, error)
}

type SeasonService interface {
	GetSeasons(SeasonFilter) (Seasons, error)
	GetSeason(string) (Season, error)
	CreateSeason(Season) (Season, error)
	DeleteSeason(string) error
}

type ChampionshipSeasonService interface {
	RegisterSeason(ChampionshipSeason) (ChampionshipSeason, error)
	GetChampionshipsSeasons(ChampionshipSeasonFilter) (ChampionshipsSeasons, error)
}

var (
	ErrClubNotFound               = errors.New("club not found")
	ErrChampionshipNotFound       = errors.New("championship not found")
	ErrSeasonNotFound             = errors.New("season not found")
	ErrSeasonAlreadyRegistered    = errors.New("season already registered for championship")
	ErrChampionshipSeasonNotFound = errors.New("ChampionshipSeason not found")
	ErrSeasonDurationInvalid      = errors.New("invalid season start and/or end values")
)
