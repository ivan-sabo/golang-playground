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

	c.GET("", gr.GetChampionships)
}

// swagger:route GET /championship Championship getChampionship
// Championship holds all data relevant to a soccer league.
//
//	Produces:
//		- application/json
//
//	responses:
//		200: GetChampionshipsResponse
func (gr *ChampionshipHandler) GetChampionships(c *gin.Context) {
	championships, err := gr.Repo.GetChampionships()
	if err != nil {
		log.Printf("An error occured: %v", err)
		return
	}

	ch := models.GetChampionshipsResponse{
		Championships: models.NewChampionshipsDTO(championships),
	}
	c.JSON(http.StatusOK, ch)
}
