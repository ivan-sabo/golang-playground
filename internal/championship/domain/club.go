package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ivan-sabo/golang-playground/internal"
)

type Clubs []Club

type Club struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewClub(id string, name string) (Club, error) {
	if id == "" {
		id = uuid.NewString()
	}

	cuuid, err := uuid.Parse(id)
	if err != nil {
		return Club{}, internal.ErrInvalidUUID
	}

	return Club{
		ID:   cuuid,
		Name: name,
	}, nil
}

type ClubFilter struct {
	Name string
}

type ClubRepo interface {
	GetClubs(context.Context, ClubFilter) (Clubs, error)
	GetClub(context.Context, string) (Club, error)
	CreateClub(context.Context, Club) (Club, error)
	UpdateClub(context.Context, string, Club) (Club, error)
	DeleteClub(context.Context, string) error
}

type ClubService interface {
	GetClubs(context.Context, ClubFilter) (Clubs, error)
	GetClub(context.Context, string) (Club, error)
	CreateClub(context.Context, Club) (Club, error)
	UpdateClub(context.Context, string, Club) (Club, error)
	DeleteClub(context.Context, string) error
}
