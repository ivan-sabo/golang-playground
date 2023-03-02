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

func (c *SeasonMySQLRepo) GetSeasons(domain.SeasonFilter) (domain.Seasons, error) {
	var seasons mysql.Seasons
	tx := c.conn.Model(&mysql.Season{}).Find(&seasons)
	if tx.Error != nil {
		return domain.Seasons{}, tx.Error
	}

	dss, err := seasons.ToEntity()
	if err != nil {
		return domain.Seasons{}, err
	}

	return dss, nil
}

func (c *SeasonMySQLRepo) GetSeason(id string) (domain.Season, error) {
	var season mysql.Season
	tx := c.conn.Where("id = ?", id).First(&season)

	if tx.Error == gorm.ErrRecordNotFound {
		return domain.Season{}, domain.ErrSeasonNotFound
	}
	if tx.Error != nil {
		return domain.Season{}, tx.Error
	}

	s, err := season.ToEntity()
	if err != nil {
		return domain.Season{}, err
	}

	return s, nil
}

func (c *SeasonMySQLRepo) CreateSeason(ds domain.Season) (domain.Season, error) {
	season := mysql.NewSeason(ds)
	tx := c.conn.Create(&season)

	if tx.Error != nil {
		return domain.Season{}, tx.Error
	}

	ds, err := season.ToEntity()
	if err != nil {
		return domain.Season{}, err
	}

	return ds, nil
}

func (c *SeasonMySQLRepo) DeleteSeason(id string) error {
	tx := c.conn.Delete(&mysql.Season{}, "id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
