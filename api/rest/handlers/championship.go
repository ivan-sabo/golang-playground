package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan-sabo/golang-playground/api/rest/models"
	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
	"github.com/ivan-sabo/golang-playground/internal/championship/infrastructure/database/mysql"
	"gorm.io/gorm"
)

type ChampionshipHandler struct {
	GinEngine *gin.Engine
	Repo      domain.ChampionshipRepo
}

func NewChampionshipHandler(ginEngine *gin.Engine, dbConn *gorm.DB) ChampionshipHandler {
	return ChampionshipHandler{
		GinEngine: ginEngine,
		Repo:      mysql.NewChampionshipMySQLRepo(dbConn),
	}
}

func (gr *ChampionshipHandler) AddChampionshipRoutes() {
	c := gr.GinEngine.Group("/championship")

	c.GET("", gr.GetChampionship)
}

// swagger:route GET /championship Championship getChampionship
// Championship holds all data relevant to a soccer league.
//
//	Produces:
//		- application/json
//
//	responses:
//		200: GetChampionshipResponse
func (gr *ChampionshipHandler) GetChampionship(c *gin.Context) {
	clubs, err := gr.Repo.GetClubs()
	if err != nil {
		log.Printf("An error occured: %v", err)
		return
	}

	ch := models.GetChampionshipResponse{
		Championship: models.ChampionshipDTO{
			Clubs: models.NewClubsDTO(clubs),
			Name:  "UEFA",
		},
	}
	c.JSON(http.StatusOK, ch)
}
