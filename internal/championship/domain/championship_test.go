package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ivan-sabo/golang-playground/internal"
)

func TestNewChampionship(t *testing.T) {
	testNewChampionshipSuccess(t)
	testNewChampionshipInvalidUUID(t)
}

func testNewChampionshipSuccess(t *testing.T) {
	name := "uefa"
	c, err := NewChampionship("", name)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.Name != name {
		t.Fatalf("expected %v, got %v", name, c.Name)
	}

	_, err = uuid.Parse(c.ID.String())
	if err != nil {
		t.Fatalf("invalid uuid: %v", err)
	}
}

func testNewChampionshipInvalidUUID(t *testing.T) {
	name := "uefa"
	invalidUUID := "test"
	c, err := NewChampionship(invalidUUID, name)

	if err != internal.ErrInvalidUUID {
		t.Fatalf("UUID should be invalid and an error should be returned, got : %v", c.ID)
	}
}

func TestNewSeason(t *testing.T) {
	testNewSeasonSuccess(t)
	testNewSeasonInvalidUUID(t)
}

func testNewSeasonSuccess(t *testing.T) {
	startYear := 2023
	endYear := 2024
	s, err := NewSeason("", startYear, endYear)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s.StartYear != startYear {
		t.Fatalf("expected %v, got %v", startYear, s.StartYear)
	}
	if s.EndYear != endYear {
		t.Fatalf("expected %v, got %v", endYear, s.EndYear)
	}
	_, err = uuid.Parse(s.ID.String())
	if err != nil {
		t.Fatalf("invalid uuid: %v", err)
	}
}

func testNewSeasonInvalidUUID(t *testing.T) {
	invalidUUID := "test"
	startYear := 2023
	endYear := 2024
	s, err := NewSeason(invalidUUID, startYear, endYear)

	if err != internal.ErrInvalidUUID {
		t.Fatalf("UUID should be invalid and an error should be returned, got : %v", s.ID)
	}
}

func TestRegisterSeason(t *testing.T) {
	cID := uuid.New()
	c := Championship{
		ID: cID,
	}
	sID := uuid.New()
	s := Season{
		ID: sID,
	}

	cs := c.RegisterSeason(s)
	if cs.Championship.ID != cID {
		t.Fatalf("expected %v, got %v", cID, cs.Championship.ID)
	}
	if cs.Season.ID != sID {
		t.Fatalf("expected %v, got %v", sID, cs.Season.ID)
	}
}

func TestChampionshipSeason(t *testing.T) {
	for text, f := range map[string]func(*testing.T, ChampionshipSeason){
		"test add club":     testChampionshipSeasonAddClub,
		"test remove club":  testChampionshipSeasonRemoveClub,
		"test add round":    testChampionshipSeasonAddRound,
		"test remove round": testChampionshipSeasonRemoveRound,
	} {
		cID := uuid.New()
		c := Championship{
			ID: cID,
		}
		sID := uuid.New()
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

	cID := uuid.New()
	c := Club{
		ID: cID,
	}

	cs.AddClub(c)
	if len(cs.Clubs) != 1 {
		t.Fatalf("there should be one club registered, have %v clubs", len(cs.Clubs))
	}
	if cs.Clubs[0].ID != cID {
		t.Fatalf("expected %v, got %v", cID, cs.Clubs[0].ID)
	}
}

func testChampionshipSeasonRemoveClub(t *testing.T, cs ChampionshipSeason) {
	t.Helper()

	cID := uuid.New()
	c := Club{
		ID: cID,
	}
	cs.Clubs = Clubs{c}
	cs.Rounds = Rounds{
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
	if len(cs.Clubs) != 0 {
		t.Fatalf("there should be no clubs registered, have %v club(s)", len(cs.Clubs))
	}
	for _, r := range cs.Rounds {
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
	if len(cs.Rounds) != 1 {
		t.Fatalf("there should be one round in ChampionshipSeason, got %v", len(cs.Rounds))
	}
	if cs.Rounds[0].NO != 1 {
		t.Fatalf("expected %v, got %v", rNO, cs.Rounds[0])
	}
}

func testChampionshipSeasonRemoveRound(t *testing.T, cs ChampionshipSeason) {
	t.Helper()

	rNO := 1
	r := Round{
		NO: rNO,
	}
	cs.Rounds = Rounds{r}
	cs.RemoveRound(r)
	if len(cs.Rounds) != 0 {
		t.Fatalf("there should be zero rounds in ChampionshipSeason, got %v", len(cs.Rounds))
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

	gID := uuid.New()
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
		ID: uuid.New(),
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
		ID: uuid.New(),
	}
	c2 := Club{
		ID: uuid.New(),
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
