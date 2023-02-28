package repository

import (
	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
	"github.com/ivan-sabo/golang-playground/internal/championship/infrastructure/database/mysql"
	"gorm.io/gorm"
)

type ChampionshipSeasonMySQLRepo struct {
	conn *gorm.DB
}

func NewChampionshipSeasonMySQLRepo(conn *gorm.DB) *ChampionshipSeasonMySQLRepo {
	return &ChampionshipSeasonMySQLRepo{
		conn: conn,
	}
}

func (r *ChampionshipSeasonMySQLRepo) RegisterSeason(ds domain.ChampionshipSeason) (domain.ChampionshipSeason, error) {
	var championship mysql.Championship
	tx := r.conn.Where("id = ?", ds.Championship.ID.String()).First(&championship)
	if tx.Error == gorm.ErrRecordNotFound {
		return domain.ChampionshipSeason{}, domain.ErrChampionshipNotFound
	}
	if tx.Error != nil {
		return domain.ChampionshipSeason{}, tx.Error
	}

	var season mysql.Season
	tx = r.conn.Where("id = ?", ds.Season.ID.String()).First(&season)
	if tx.Error == gorm.ErrRecordNotFound {
		return domain.ChampionshipSeason{}, domain.ErrSeasonNotFound
	}
	if tx.Error != nil {
		return domain.ChampionshipSeason{}, tx.Error
	}

	var championshipsSeasons mysql.ChampionshipsSeasons
	tx = r.conn.Where(
		"championship_id = ? AND season_id = ?",
		ds.Championship.ID.String(),
		ds.Season.ID.String(),
	).Find(&championshipsSeasons)
	if tx.Error != nil {
		return domain.ChampionshipSeason{}, tx.Error
	}
	if len(championshipsSeasons) > 0 {
		return domain.ChampionshipSeason{}, domain.ErrSeasonAlreadyRegistered
	}

	championshipSeason := mysql.NewChampionshipSeason(ds.Championship, ds.Season)
	tx = r.conn.Create(&championshipSeason)
	if tx.Error != nil {
		return domain.ChampionshipSeason{}, tx.Error
	}

	dc, err := championship.ToEntity()
	if err != nil {
		return domain.ChampionshipSeason{}, err
	}
	dse, err := season.ToEntity()
	if err != nil {
		return domain.ChampionshipSeason{}, err
	}

	return domain.ChampionshipSeason{
		Championship: dc,
		Season:       dse,
	}, nil
}

func (r *ChampionshipSeasonMySQLRepo) GetChampionshipsSeasons(csf domain.ChampionshipSeasonFilter) (domain.ChampionshipsSeasons, error) {
	var championshipsSeasons mysql.ChampionshipsSeasons
	tx := r.conn.Where(
		"championship_id = ? AND season_id = ?",
		csf.ChampionshipID,
		csf.SeasonID,
	).First(&championshipsSeasons)
	if tx.Error == gorm.ErrRecordNotFound {
		return domain.ChampionshipsSeasons{}, domain.ErrChampionshipSeasonNotFound
	}
	if tx.Error != nil {
		return domain.ChampionshipsSeasons{}, tx.Error
	}

	csIDs, ssIDs := getChampionshipAndSeasonIDs(championshipsSeasons)

	// @todo: use eager loading to reduce number of queries
	var championships mysql.Championships
	tx = r.conn.Where("id IN ?", csIDs).Find(&championships)
	if tx.Error != nil {
		return domain.ChampionshipsSeasons{}, tx.Error
	}

	var seasons mysql.Seasons
	tx = r.conn.Where("id IN ?", ssIDs).Find(&seasons)
	if tx.Error != nil {
		return domain.ChampionshipsSeasons{}, tx.Error
	}

	dcs, err := championships.ToEntity()
	if err != nil {
		return domain.ChampionshipsSeasons{}, err
	}
	dss, err := seasons.ToEntity()
	if err != nil {
		return domain.ChampionshipsSeasons{}, err
	}

	return championshipsSeasons.ToEntity(dcs, dss), nil
}

func getChampionshipAndSeasonIDs(csss mysql.ChampionshipsSeasons) ([]string, []string) {
	championshipsIDs := map[string]bool{}
	seasonsIDs := map[string]bool{}
	for _, csss := range csss {
		if _, ok := championshipsIDs[csss.ChampionshipID]; !ok {
			championshipsIDs[csss.ChampionshipID] = true
		}

		if _, ok := seasonsIDs[csss.SeasonID]; !ok {
			seasonsIDs[csss.SeasonID] = true
		}
	}

	csIDs := make([]string, 0, len(championshipsIDs))
	for cID, _ := range championshipsIDs {
		csIDs = append(csIDs, cID)
	}

	ssIDs := make([]string, 0, len(seasonsIDs))
	for sID, _ := range seasonsIDs {
		ssIDs = append(ssIDs, sID)
	}

	return csIDs, ssIDs
}
