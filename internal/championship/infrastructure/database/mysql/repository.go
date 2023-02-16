package mysql

import (
	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
	"gorm.io/gorm"
)

type ChampionshipRepo interface {
	GetClubs() (domain.Clubs, error)
}

type ChampionshipMySQLRepo struct {
	conn *gorm.DB
}

func NewChampionshipMySQLRepo(conn *gorm.DB) *ChampionshipMySQLRepo {
	return &ChampionshipMySQLRepo{
		conn: conn,
	}
}

func (r ChampionshipMySQLRepo) GetClubs() (domain.Clubs, error) {
	var clubs ClubsGorm
	tx := r.conn.Limit(10).Find(clubs)

	if tx.Error != nil {
		return domain.Clubs{}, tx.Error
	}

	return clubs.ToEntity(), nil
}
