package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan-sabo/golang-playground/api/rest/models"
	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
	"github.com/ivan-sabo/golang-playground/internal/championship/infrastructure/database/mysql/repository"
	"gorm.io/gorm"
)

type SeasonHandler struct {
	GinEngine *gin.Engine
	Repo      domain.SeasonRepo
}

func NewSeasonHandler(g *gin.Engine, conn *gorm.DB) SeasonHandler {
	return SeasonHandler{
		GinEngine: g,
		Repo:      repository.NewSeasonMySQLRepo(conn),
	}
}

func (h *SeasonHandler) AddSeasonRoutes() {
	g := h.GinEngine.Group("/season")

	g.GET("", h.getSeasons)
	g.GET("/:id", h.getSeason)
	g.POST("", h.createSeason)
	g.PUT("", h.updateSeason)
	g.DELETE("/:id", h.deleteSeason)
}

// swagger:route GET /season Season getSeasons
// Get a list of seasons.
//
//	Produces:
//		- application/json
//
//	responses:
//		200: GetSeasonsResponse
func (h *SeasonHandler) getSeasons(c *gin.Context) {
	seasons, err := h.Repo.GetSeasons(domain.SeasonFilter{})
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusNotFound, models.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.NewGetSeasonsResponse(seasons))
}

func (h *SeasonHandler) getSeason(c *gin.Context) {

}

// swagger:route POST /season Season CreateSeason
// Create a new Season.
//
//	Consumes:
//		- application/json
//
//	Produces:
//		- application/json
//
//	responses:
//		200: PostSeasonResponse
//		500: ErrorResponse
func (h *SeasonHandler) createSeason(c *gin.Context) {
	var season models.PostSeasonRequest
	err := c.ShouldBindJSON(&season)
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	ds, err := h.Repo.CreateSeason(season.ToEntity())
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, models.NewPostSeasonResponse(ds))
}

func (h *SeasonHandler) updateSeason(c *gin.Context) {

}

func (h *SeasonHandler) deleteSeason(c *gin.Context) {

}
