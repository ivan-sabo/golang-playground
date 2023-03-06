package repository

import (
	"context"

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

func (c *SeasonMySQLRepo) GetSeasons(ctx context.Context, f domain.SeasonFilter) (domain.Seasons, error) {
	var seasons mysql.Seasons
	tx := c.conn.WithContext(ctx).Model(&mysql.Season{}).Find(&seasons)
	if tx.Error != nil {
		return domain.Seasons{}, tx.Error
	}

	dss, err := seasons.ToEntity()
	if err != nil {
		return domain.Seasons{}, err
	}

	return dss, nil
}

func (c *SeasonMySQLRepo) GetSeason(ctx context.Context, id string) (domain.Season, error) {
	var season mysql.Season
	tx := c.conn.WithContext(ctx).Where("id = ?", id).First(&season)

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

func (c *SeasonMySQLRepo) CreateSeason(ctx context.Context, ds domain.Season) (domain.Season, error) {
	season := mysql.NewSeason(ds)
	tx := c.conn.WithContext(ctx).Create(&season)

	if tx.Error != nil {
		return domain.Season{}, tx.Error
	}

	ds, err := season.ToEntity()
	if err != nil {
		return domain.Season{}, err
	}

	return ds, nil
}

func (c *SeasonMySQLRepo) DeleteSeason(ctx context.Context, id string) error {
	tx := c.conn.WithContext(ctx).Delete(&mysql.Season{}, "id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
