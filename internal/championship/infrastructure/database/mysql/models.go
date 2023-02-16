package mysql

import (
	"time"

	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
	"gorm.io/gorm"
)

type ClubsGorm []ClubGorm

type ClubGorm struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (c *ClubGorm) ToEntity() domain.Club {
	return domain.Club{
		ID:   c.ID,
		Name: c.Name,
	}
}

func (cs ClubsGorm) ToEntity() domain.Clubs {
	clubs := make(domain.Clubs, len(cs))
	for _, c := range cs {
		clubs = append(clubs, c.ToEntity())
	}

	return clubs
}
