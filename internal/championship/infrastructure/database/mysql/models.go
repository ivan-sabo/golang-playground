package mysql

import (
	"time"

	"github.com/google/uuid"
	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
	"gorm.io/gorm"
)

type Clubs []Club

type Club struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func NewClub(c domain.Club) Club {
	return Club{
		ID:   uuid.NewString(),
		Name: c.Name,
	}
}

func (c *Club) ToEntity() domain.Club {
	return domain.Club{
		ID:   c.ID,
		Name: c.Name,
	}
}

func (cs Clubs) ToEntity() domain.Clubs {
	clubs := make(domain.Clubs, 0, len(cs))
	for _, c := range cs {
		clubs = append(clubs, c.ToEntity())
	}

	return clubs
}

type Championships []Championship

type Championship struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (c *Championship) ToEntity() domain.Championship {
	return domain.Championship{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		DeletedAt: c.DeletedAt.Time,
	}
}

func (cs Championships) ToEntity() domain.Championships {
	championships := make(domain.Championships, 0, len(cs))
	for _, c := range cs {
		championships = append(championships, c.ToEntity())
	}

	return championships
}
