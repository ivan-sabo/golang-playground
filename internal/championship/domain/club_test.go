package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ivan-sabo/golang-playground/internal"
)

func TestNewClub(t *testing.T) {
	testNewClubSuccess(t)
	testNewClubInvalidID(t)
}

func testNewClubSuccess(t *testing.T) {
	n := "borussia"
	c, err := NewClub("", n)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.Name != n {
		t.Fatalf("expected %v, got %v", n, c.Name)
	}

	_, err = uuid.Parse(c.ID.String())
	if err != nil {
		t.Fatalf("invalid uuid: %v", err)
	}
}

func testNewClubInvalidID(t *testing.T) {
	n := "borussia"
	invalidID := "test"
	c, err := NewClub(invalidID, n)

	if err != internal.ErrInvalidUUID {
		t.Fatalf("UUID should be invalid and an error should be returned, got : %v", c.ID)
	}
}
