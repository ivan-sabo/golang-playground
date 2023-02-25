package domain

import "testing"

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
