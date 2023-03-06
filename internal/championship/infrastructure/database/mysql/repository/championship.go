package repository

import (
	"context"

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

func (r ChampionshipMySQLRepo) GetChampionships(ctx context.Context, df domain.ChampionshipFilter) (domain.Championships, error) {
	// @todo: implement filter
	var championships mysql.Championships
	tx := r.conn.WithContext(ctx).Model(&mysql.Championship{}).Find(&championships)
	if tx.Error != nil {
		return domain.Championships{}, tx.Error
	}

	dcs, err := championships.ToEntity()
	if err != nil {
		return domain.Championships{}, err
	}

	return dcs, nil
}

func (r ChampionshipMySQLRepo) GetChampionship(ctx context.Context, id string) (domain.Championship, error) {
	var championship mysql.Championship
	tx := r.conn.WithContext(ctx).Where("id = ?", id).First(&championship)

	if tx.Error == gorm.ErrRecordNotFound {
		return domain.Championship{}, domain.ErrChampionshipNotFound
	}
	if tx.Error != nil {
		return domain.Championship{}, tx.Error
	}

	dc, err := championship.ToEntity()
	if err != nil {
		return domain.Championship{}, err
	}

	return dc, nil
}

func (c ChampionshipMySQLRepo) CreateChampionship(ctx context.Context, dc domain.Championship) (domain.Championship, error) {
	championship := mysql.NewChampionship(dc)
	tx := c.conn.WithContext(ctx).Create(&championship)

	if tx.Error != nil {
		return domain.Championship{}, tx.Error
	}

	dc, err := championship.ToEntity()
	if err != nil {
		return domain.Championship{}, err
	}

	return dc, nil
}

func (c ChampionshipMySQLRepo) UpdateChampionship(ctx context.Context, id string, dc domain.Championship) (domain.Championship, error) {
	var championship mysql.Championship
	tx := c.conn.WithContext(ctx).Where("id = ?", id).First(&championship)

	if tx.Error == gorm.ErrRecordNotFound {
		return domain.Championship{}, domain.ErrChampionshipNotFound
	}
	if tx.Error != nil {
		return domain.Championship{}, tx.Error
	}

	championship.Name = dc.Name
	tx.Save(&championship)
	if tx.Error != nil {
		return domain.Championship{}, tx.Error
	}

	dc, err := championship.ToEntity()
	if err != nil {
		return domain.Championship{}, err
	}

	return dc, nil
}

func (c ChampionshipMySQLRepo) DeleteChampionship(ctx context.Context, id string) error {
	tx := c.conn.WithContext(ctx).Delete(&mysql.Championship{}, "id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
