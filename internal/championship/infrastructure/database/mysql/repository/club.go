package repository

import (
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

func (c ClubMySQLRepo) GetClubs(domain.ClubFilter) (domain.Clubs, error) {
	var clubs mysql.Clubs
	tx := c.conn.Model(&mysql.Club{}).Find(&clubs)
	if tx.Error != nil {
		return domain.Clubs{}, tx.Error
	}

	dcs, err := clubs.ToEntity()
	if err != nil {
		return domain.Clubs{}, err
	}

	return dcs, nil
}

func (c ClubMySQLRepo) GetClub(id string) (domain.Club, error) {
	var club mysql.Club
	tx := c.conn.Where("id = ?", id).First(&club)

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

func (c ClubMySQLRepo) CreateClub(dc domain.Club) (domain.Club, error) {
	club := mysql.NewClub(dc)
	tx := c.conn.Create(&club)

	if tx.Error != nil {
		return domain.Club{}, tx.Error
	}

	dc, err := club.ToEntity()
	if err != nil {
		return domain.Club{}, err
	}

	return dc, nil
}

func (c ClubMySQLRepo) UpdateClub(id string, dc domain.Club) (domain.Club, error) {
	var club mysql.Club
	tx := c.conn.Where("id = ?", id).First(&club)

	if tx.Error == gorm.ErrRecordNotFound {
		return domain.Club{}, domain.ErrClubNotFound
	}
	if tx.Error != nil {
		return domain.Club{}, tx.Error
	}

	club.Name = dc.Name
	tx = c.conn.Save(&club)
	if tx.Error != nil {
		return domain.Club{}, tx.Error
	}

	dc, err := club.ToEntity()
	if err != nil {
		return domain.Club{}, err
	}

	return dc, nil
}

func (c ClubMySQLRepo) DeleteClub(id string) error {
	tx := c.conn.Delete(&mysql.Club{}, "id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
