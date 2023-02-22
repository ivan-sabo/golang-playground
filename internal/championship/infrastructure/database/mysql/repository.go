package mysql

import (
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

func (r ChampionshipMySQLRepo) GetChampionships() (domain.Championships, error) {
	var championships Championships
	tx := r.conn.Model(&Championship{}).Find(&championships)

	if tx.Error != nil {
		return domain.Championships{}, tx.Error
	}

	return championships.ToEntity(), nil
}
