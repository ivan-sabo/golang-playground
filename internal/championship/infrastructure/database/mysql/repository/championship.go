package repository

import (
	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
	"github.com/ivan-sabo/golang-playground/internal/championship/infrastructure/database/mysql"
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
	var championships mysql.Championships
	tx := r.conn.Model(&mysql.Championship{}).Find(&championships)
	if tx.Error != nil {
		return domain.Championships{}, tx.Error
	}

	return championships.ToEntity(), nil
}

func (r ChampionshipMySQLRepo) GetChampionship(id string) (domain.Championship, error) {
	var championship mysql.Championship
	tx := r.conn.Where("id = ?", id).First(&championship)

	if tx.Error == gorm.ErrRecordNotFound {
		return domain.Championship{}, domain.ErrChampionshipNotFound
	}
	if tx.Error != nil {
		return domain.Championship{}, tx.Error
	}

	return championship.ToEntity(), nil
}

func (c ChampionshipMySQLRepo) CreateChampionship(dc domain.Championship) (domain.Championship, error) {
	championship := mysql.NewChampionship(dc)
	tx := c.conn.Create(&championship)

	if tx.Error != nil {
		return domain.Championship{}, tx.Error
	}

	return championship.ToEntity(), nil
}
