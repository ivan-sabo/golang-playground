package mysql

import (
	"time"

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
