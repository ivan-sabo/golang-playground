package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/ivan-sabo/golang-playground/internal"
)

var (
	Created    Status = "created"
	Scheduled  Status = "scheduled"
	InProgress Status = "in progress"
	Finished   Status = "finished"
)

type Games []Game

type Game struct {
	ID        uuid.UUID
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

func NewGame(id string, home Club, away Club) (Game, error) {
	if id == "" {
		id = uuid.NewString()
	}

	cuuid, err := uuid.Parse(id)
	if err != nil {
		return Game{}, internal.ErrInvalidUUID
	}

	return Game{
		ID:     cuuid,
		Home:   home,
		Away:   away,
		Status: Created,
	}, nil
}
