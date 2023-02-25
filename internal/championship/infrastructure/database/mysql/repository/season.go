package repository

import (
	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
	"github.com/ivan-sabo/golang-playground/internal/championship/infrastructure/database/mysql"
	"gorm.io/gorm"
)

type SeasonMySQLRepo struct {
	conn *gorm.DB
}

func NewSeasonMySQLRepo(conn *gorm.DB) *SeasonMySQLRepo {
	return &SeasonMySQLRepo{
		conn: conn,
	}
}

func (c SeasonMySQLRepo) GetSeasons(domain.SeasonFilter) (domain.Seasons, error) {
	var seasons mysql.Seasons
	tx := c.conn.Model(&mysql.Season{}).Find(&seasons)
	if tx.Error != nil {
		return domain.Seasons{}, tx.Error
	}

	return seasons.ToEntity(), nil
}
