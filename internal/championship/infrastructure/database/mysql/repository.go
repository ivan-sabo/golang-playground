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

type ClubMySQLRepo struct {
	conn *gorm.DB
}

func NewClubMySQLRepo(conn *gorm.DB) *ClubMySQLRepo {
	return &ClubMySQLRepo{
		conn: conn,
	}
}

func (c ClubMySQLRepo) GetClubs(domain.ClubFilter) (domain.Clubs, error) {
	var clubs Clubs
	tx := c.conn.Model(&Club{}).Find(&clubs)
	if tx.Error != nil {
		return domain.Clubs{}, tx.Error
	}

	return clubs.ToEntity(), nil
}

func (c ClubMySQLRepo) GetClub(id string) (domain.Club, error) {
	var club Club
	tx := c.conn.Where("id = ?", id).First(&club)

	if tx.Error == gorm.ErrRecordNotFound {
		return domain.Club{}, domain.ErrClubNotFound
	}
	if tx.Error != nil {
		return domain.Club{}, tx.Error
	}

	return club.ToEntity(), nil
}

func (c ClubMySQLRepo) CreateClub(dc domain.Club) (domain.Club, error) {
	club := NewClub(dc)
	tx := c.conn.Create(&club)

	if tx.Error != nil {
		return domain.Club{}, tx.Error
	}

	return club.ToEntity(), nil
}

func (c ClubMySQLRepo) UpdateClub(id string, dc domain.Club) (domain.Club, error) {
	var club Club
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

	return club.ToEntity(), nil
}
