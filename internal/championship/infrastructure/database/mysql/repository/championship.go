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

func (r ChampionshipMySQLRepo) GetChampionships(df domain.ChampionshipFilter) (domain.Championships, error) {
	// @todo: implement filter
	var championships mysql.Championships
	tx := r.conn.Model(&mysql.Championship{}).Find(&championships)
	if tx.Error != nil {
		return domain.Championships{}, tx.Error
	}

	dcs, err := championships.ToEntity()
	if err != nil {
		return domain.Championships{}, err
	}

	return dcs, nil
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

	dc, err := championship.ToEntity()
	if err != nil {
		return domain.Championship{}, err
	}

	return dc, nil
}

func (c ChampionshipMySQLRepo) CreateChampionship(dc domain.Championship) (domain.Championship, error) {
	championship := mysql.NewChampionship(dc)
	tx := c.conn.Create(&championship)

	if tx.Error != nil {
		return domain.Championship{}, tx.Error
	}

	dc, err := championship.ToEntity()
	if err != nil {
		return domain.Championship{}, err
	}

	return dc, nil
}

func (c ChampionshipMySQLRepo) UpdateChampionship(id string, dc domain.Championship) (domain.Championship, error) {
	var championship mysql.Championship
	tx := c.conn.Where("id = ?", id).First(&championship)

	if tx.Error == gorm.ErrRecordNotFound {
		return domain.Championship{}, domain.ErrChampionshipNotFound
	}
	if tx.Error != nil {
		return domain.Championship{}, tx.Error
	}

	championship.Name = dc.Name
	tx = c.conn.Save(&championship)
	if tx.Error != nil {
		return domain.Championship{}, tx.Error
	}

	dc, err := championship.ToEntity()
	if err != nil {
		return domain.Championship{}, err
	}

	return dc, nil
}

func (c ChampionshipMySQLRepo) DeleteChampionship(id string) error {
	tx := c.conn.Delete(&mysql.Championship{}, "id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
