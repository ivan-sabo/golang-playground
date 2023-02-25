package domain

import (
	"time"

	"github.com/google/uuid"
)

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
