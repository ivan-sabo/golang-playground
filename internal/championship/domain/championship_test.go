package domain

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNewChampionship(t *testing.T) {
	name := "uefa"
	c := NewChampionship(name)

	if c.Name != name {
		t.Fatalf("expected %v, got %v", name, c.Name)
	}
	if len(c.ID) != 36 {
		t.Fatalf("UUID should be 36 char long, current value: %v", c.ID)
	}
}

func TestNewSeason(t *testing.T) {
	startYear := 2023
	endYear := 2024
	s := NewSeason(startYear, endYear)

	if s.StartYear != startYear {
		t.Fatalf("expected %v, got %v", startYear, s.StartYear)
	}
	if s.EndYear != endYear {
		t.Fatalf("expected %v, got %v", endYear, s.EndYear)
	}
	if len(s.ID) != 36 {
		t.Fatalf("UUID should be 36 char long, current value: %v", s.ID)
	}
}

func TestRegisterSeason(t *testing.T) {
	cID := uuid.NewString()
	c := Championship{
		ID: cID,
	}
	sID := uuid.NewString()
	s := Season{
		ID: sID,
	}

	cs := c.RegisterSeason(s)
	if cs.championship.ID != cID {
		t.Fatalf("expected %v, got %v", cID, cs.championship.ID)
	}
	if cs.season.ID != sID {
		t.Fatalf("expected %v, got %v", sID, cs.season.ID)
	}
}

func TestNewClub(t *testing.T) {
	n := "borussia"
	c := NewClub(n)

	if c.Name != n {
		t.Fatalf("expected %v, got %v", n, c.Name)
	}
	if len(c.ID) != 36 {
		t.Fatalf("UUID should be 36 char long, current value: %v", c.ID)
	}
}

func TestChampionshipSeason(t *testing.T) {
	for text, f := range map[string]func(*testing.T, ChampionshipSeason){
		"test add club":     testChampionshipSeasonAddClub,
		"test remove club":  testChampionshipSeasonRemoveClub,
		"test add round":    testChampionshipSeasonAddRound,
		"test remove round": testChampionshipSeasonRemoveRound,
	} {
		cID := uuid.NewString()
		c := Championship{
			ID: cID,
		}
		sID := uuid.NewString()
		s := Season{
			ID: sID,
		}
		cs := c.RegisterSeason(s)

		t.Log(text)
		f(t, cs)
	}
}

func testChampionshipSeasonAddClub(t *testing.T, cs ChampionshipSeason) {
	t.Helper()

	cID := uuid.NewString()
	c := Club{
		ID: cID,
	}

	cs.AddClub(c)
	if len(cs.clubs) != 1 {
		t.Fatalf("there should be one club registered, have %v clubs", len(cs.clubs))
	}
	if cs.clubs[0].ID != cID {
		t.Fatalf("expected %v, got %v", cID, cs.clubs[0].ID)
	}
}

func testChampionshipSeasonRemoveClub(t *testing.T, cs ChampionshipSeason) {
	t.Helper()

	cID := uuid.NewString()
	c := Club{
		ID: cID,
	}
	cs.clubs = Clubs{c}
	cs.rounds = Rounds{
		{
			NO: 1,
			Games: Games{
				{
					Home: c,
				},
			},
		},
		{
			NO: 2,
			Games: Games{
				{
					Away: c,
				},
			},
		},
	}
	cs.RemoveClub(c)
	if len(cs.clubs) != 0 {
		t.Fatalf("there should be no clubs registered, have %v club(s)", len(cs.clubs))
	}
	for _, r := range cs.rounds {
		if len(r.Games) != 0 {
			t.Fatalf("all rounds should be empty, got %#v", r.Games)
		}
	}
}

func testChampionshipSeasonAddRound(t *testing.T, cs ChampionshipSeason) {
	t.Helper()

	rNO := 1
	r := Round{
		NO: rNO,
	}
	cs.AddRound(r)
	if len(cs.rounds) != 1 {
		t.Fatalf("there should be one round in ChampionshipSeason, got %v", len(cs.rounds))
	}
	if cs.rounds[0].NO != 1 {
		t.Fatalf("expected %v, got %v", rNO, cs.rounds[0])
	}
}

func testChampionshipSeasonRemoveRound(t *testing.T, cs ChampionshipSeason) {
	t.Helper()

	rNO := 1
	r := Round{
		NO: rNO,
	}
	cs.rounds = Rounds{r}
	cs.RemoveRound(r)
	if len(cs.rounds) != 0 {
		t.Fatalf("there should be zero rounds in ChampionshipSeason, got %v", len(cs.rounds))
	}
}

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

func TestRound(t *testing.T) {
	for text, f := range map[string]func(*testing.T, Round){
		"test add game":    testRoundAddGame,
		"test remove game": testRoundRemoveGame,
		"test remove club": testRoundRemoveClub,
	} {
		r := Round{
			NO: 1,
		}

		t.Logf(text)
		f(t, r)
	}
}

func testRoundAddGame(t *testing.T, r Round) {
	t.Helper()

	gID := uuid.NewString()
	g := Game{
		ID: gID,
	}
	r.AddGame(g)
	if len(r.Games) != 1 {
		t.Fatalf("there should be one game, got %v", len(r.Games))
	}
	if r.Games[0].ID != gID {
		t.Fatalf("expected %v, got %v", gID, r.Games[0].ID)
	}
}

func testRoundRemoveGame(t *testing.T, r Round) {
	t.Helper()

	g := Game{
		ID: uuid.NewString(),
	}
	r.Games = Games{g}
	r.RemoveGame(g)
	if len(r.Games) != 0 {
		t.Fatalf("there should be no games in round, got %v", len(r.Games))
	}
}

func testRoundRemoveClub(t *testing.T, r Round) {
	t.Helper()

	c := Club{
		ID: uuid.NewString(),
	}
	c2 := Club{
		ID: uuid.NewString(),
	}

	r.Games = Games{
		{
			Home: c,
		},
		{
			Away: c,
		},
		{
			Home: c2,
		},
	}
	r.removeClub(c)
	if len(r.Games) != 1 {
		t.Fatalf("there should be one game, got %#v", r.Games)
	}
	if r.Games[0].Home.ID != c2.ID {
		t.Fatalf("expected %v, got %#v", c2.ID, r.Games[0].Home.ID)
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
