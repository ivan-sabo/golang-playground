package domain

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNewGame(t *testing.T) {
	hc := Club{
		ID: uuid.NewString(),
	}
	ac := Club{
		ID: uuid.NewString(),
	}
	g := NewGame(hc, ac)

	if g.Home.ID != hc.ID {
		t.Fatalf("expected %v, got %v", hc.ID, g.Home.ID)
	}
	if g.Away.ID != ac.ID {
		t.Fatalf("expected %v, got %v", ac.ID, g.Away.ID)
	}
	if g.Status != Created {
		t.Fatalf("expected %v, got %v", Created, g.Status)
	}
}

func TestGame(t *testing.T) {
	for text, f := range map[string]func(*testing.T, Game){
		"test finish game":   testFinish,
		"test schedule game": testSchedule,
		"test update score":  testUpdateScore,
	} {
		g := Game{
			Home: Club{
				ID: uuid.NewString(),
			},
			Away: Club{
				ID: uuid.NewString(),
			},
			Status: Created,
		}
		t.Log(text)
		f(t, g)
	}
}

func testFinish(t *testing.T, g Game) {
	t.Helper()

	g.Finish()
	if g.Status != Finished {
		t.Fatalf("expected %v, got %v", Finished, g.Status)
	}
}

func testSchedule(t *testing.T, g Game) {
	t.Helper()

	st := time.Now()
	g.Schedule(st)
	if g.Start != st {
		t.Fatalf("expected %v, got %v", st, g.Start)
	}
	if g.Status != Scheduled {
		t.Fatalf("expected %v, got %v", Scheduled, g.Status)
	}
}

func testUpdateScore(t *testing.T, g Game) {
	t.Helper()

	g.UpdateScore(2, 3)
	if g.AwayScore != 3 {
		t.Fatalf("expected %v, got %v", 3, g.AwayScore)
	}
	if g.HomeScore != 2 {
		t.Fatalf("expected %v, got %v", 2, g.HomeScore)
	}
}
