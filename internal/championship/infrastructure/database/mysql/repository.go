package mysql

import (
	"log"

	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
	"gorm.io/gorm"
)

type ChampionshipMySQLRepo struct {
	conn *gorm.DB
}

func NewChampionshipMySQLRepo(conn *gorm.DB) *ChampionshipMySQLRepo {
	return &ChampionshipMySQLRepo{
		conn: conn,
	}
}

func (r ChampionshipMySQLRepo) GetClubs() (domain.Clubs, error) {
	var clubs Clubs
	tx := r.conn.Model(&Club{}).Find(&clubs)

	if tx.Error != nil {
		return domain.Clubs{}, tx.Error
	}

	log.Printf("got %d clubs", len(clubs))

	return clubs.ToEntity(), nil
}
