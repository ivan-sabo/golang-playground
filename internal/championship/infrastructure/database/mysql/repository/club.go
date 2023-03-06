package repository

import (
	"context"

	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
	"github.com/ivan-sabo/golang-playground/internal/championship/infrastructure/database/mysql"
	"gorm.io/gorm"
)

type ClubMySQLRepo struct {
	conn *gorm.DB
}

func NewClubMySQLRepo(conn *gorm.DB) *ClubMySQLRepo {
	return &ClubMySQLRepo{
		conn: conn,
	}
}

func (c ClubMySQLRepo) GetClubs(ctx context.Context, f domain.ClubFilter) (domain.Clubs, error) {
	// @todo: use filter
	var clubs mysql.Clubs
	tx := c.conn.WithContext(ctx).Model(&mysql.Club{}).Find(&clubs)
	if tx.Error != nil {
		return domain.Clubs{}, tx.Error
	}

	dcs, err := clubs.ToEntity()
	if err != nil {
		return domain.Clubs{}, err
	}

	return dcs, nil
}

func (c ClubMySQLRepo) GetClub(ctx context.Context, id string) (domain.Club, error) {
	var club mysql.Club
	tx := c.conn.WithContext(ctx).Where("id = ?", id).First(&club)

	if tx.Error == gorm.ErrRecordNotFound {
		return domain.Club{}, domain.ErrClubNotFound
	}
	if tx.Error != nil {
		return domain.Club{}, tx.Error
	}

	sc, err := club.ToEntity()
	if err != nil {
		return domain.Club{}, err
	}

	return sc, nil
}

func (c ClubMySQLRepo) CreateClub(ctx context.Context, dc domain.Club) (domain.Club, error) {
	club := mysql.NewClub(dc)
	tx := c.conn.WithContext(ctx).Create(&club)

	if tx.Error != nil {
		return domain.Club{}, tx.Error
	}

	dc, err := club.ToEntity()
	if err != nil {
		return domain.Club{}, err
	}

	return dc, nil
}

func (c ClubMySQLRepo) UpdateClub(ctx context.Context, id string, dc domain.Club) (domain.Club, error) {
	var club mysql.Club
	tx := c.conn.WithContext(ctx).Where("id = ?", id).First(&club)

	if tx.Error == gorm.ErrRecordNotFound {
		return domain.Club{}, domain.ErrClubNotFound
	}
	if tx.Error != nil {
		return domain.Club{}, tx.Error
	}

	club.Name = dc.Name
	tx.Save(&club)
	if tx.Error != nil {
		return domain.Club{}, tx.Error
	}

	dc, err := club.ToEntity()
	if err != nil {
		return domain.Club{}, err
	}

	return dc, nil
}

func (c ClubMySQLRepo) DeleteClub(ctx context.Context, id string) error {
	tx := c.conn.WithContext(ctx).Delete(&mysql.Club{}, "id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
