package mysql

import (
	"time"

	"github.com/google/uuid"
	"github.com/ivan-sabo/golang-playground/internal"
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
		ID:   c.ID.String(),
		Name: c.Name,
	}
}

func (c *Club) ToEntity() (domain.Club, error) {
	id, err := uuid.Parse(c.ID)
	if err != nil {
		return domain.Club{}, internal.ErrInvalidUUID
	}

	return domain.Club{
		ID:   id,
		Name: c.Name,
	}, nil
}

func (cs Clubs) ToEntity() (domain.Clubs, error) {
	clubs := make(domain.Clubs, 0, len(cs))

	for _, c := range cs {
		dc, err := c.ToEntity()
		if err != nil {
			return domain.Clubs{}, err
		}

		clubs = append(clubs, dc)
	}

	return clubs, nil
}

type Championships []Championship

type Championship struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Seasons   Seasons `gorm:"many2many:championship_season;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func NewChampionship(c domain.Championship) Championship {
	return Championship{
		ID:   c.ID.String(),
		Name: c.Name,
	}
}

func (c *Championship) ToEntity() (domain.Championship, error) {
	id, err := uuid.Parse(c.ID)
	if err != nil {
		return domain.Championship{}, internal.ErrInvalidUUID
	}

	return domain.Championship{
		ID:        id,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		DeletedAt: c.DeletedAt.Time,
	}, nil
}

func (cs Championships) ToEntity() (domain.Championships, error) {
	championships := make(domain.Championships, 0, len(cs))
	for _, c := range cs {
		dc, err := c.ToEntity()
		if err != nil {
			return domain.Championships{}, err
		}

		championships = append(championships, dc)
	}

	return championships, nil
}

type Seasons []Season

type Season struct {
	ID        string `gorm:"primaryKey"`
	StartYear int
	EndYear   int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func NewSeason(ds domain.Season) Season {
	return Season{
		ID:        ds.ID.String(),
		StartYear: ds.StartYear,
		EndYear:   ds.EndYear,
	}
}

func (ss Seasons) ToEntity() (domain.Seasons, error) {
	seasons := make(domain.Seasons, 0, len(ss))
	for _, s := range ss {
		ds, err := s.ToEntity()
		if err != nil {
			return domain.Seasons{}, err
		}
		seasons = append(seasons, ds)
	}

	return seasons, nil
}

func (s Season) ToEntity() (domain.Season, error) {
	id, err := uuid.Parse(s.ID)
	if err != nil {
		return domain.Season{}, internal.ErrInvalidUUID
	}

	return domain.Season{
		ID:        id,
		StartYear: s.StartYear,
		EndYear:   s.EndYear,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
		DeletedAt: s.DeletedAt.Time,
	}, nil
}

type ChampionshipsSeasons []ChampionshipSeason

type ChampionshipSeason struct {
	ChampionshipID string `gorm:"primaryKey"`
	SeasonID       string `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

func (cs ChampionshipSeason) TableName() string {
	return "championship_season"
}

func (cs *ChampionshipSeason) ToEntity(dc domain.Championship, ds domain.Season) domain.ChampionshipSeason {
	return domain.ChampionshipSeason{
		Championship: dc,
		Season:       ds,
		CreatedAt:    cs.CreatedAt,
		UpdatedAt:    cs.UpdatedAt,
		DeletedAt:    cs.DeletedAt.Time,
	}
}

func NewChampionshipSeason(c domain.Championship, s domain.Season) ChampionshipSeason {
	return ChampionshipSeason{
		ChampionshipID: c.ID.String(),
		SeasonID:       s.ID.String(),
	}
}

func (cs ChampionshipsSeasons) ToEntity(dcs domain.Championships, dss domain.Seasons) domain.ChampionshipsSeasons {
	championshipsSeasons := make(domain.ChampionshipsSeasons, 0, len(cs))
	mdcs := mapChampionships(dcs)
	mdss := mapSeasons(dss)

	for _, c := range cs {
		championshipsSeasons = append(championshipsSeasons, c.ToEntity(mdcs[c.ChampionshipID], mdss[c.SeasonID]))
	}

	return championshipsSeasons
}

func mapChampionships(dcs domain.Championships) map[string]domain.Championship {
	m := make(map[string]domain.Championship)
	for _, dc := range dcs {
		m[dc.ID.String()] = dc
	}

	return m
}

func mapSeasons(dss domain.Seasons) map[string]domain.Season {
	m := make(map[string]domain.Season)
	for _, ds := range dss {
		m[ds.ID.String()] = ds
	}

	return m
}
