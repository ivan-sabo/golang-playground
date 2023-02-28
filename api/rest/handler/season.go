package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/ivan-sabo/golang-playground/api/rest/model"
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
	g := h.GinEngine.Group("/seasons")

	g.GET("", h.getSeasons)
	g.GET("/:id", h.getSeason)
	g.POST("", h.createSeason)
	g.PUT("", h.updateSeason)
	g.DELETE("/:id", h.deleteSeason)
}

// swagger:route GET /seasons Season getSeasons
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

// swagger:route POST /seasons Season CreateSeason
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
	var sr models.PostSeasonRequest
	err := c.ShouldBindJSON(&sr)
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	season, err := sr.ToEntity()
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	ds, err := h.Repo.CreateSeason(season)
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, models.NewPostSeasonResponse(ds))
}

func (h *SeasonHandler) updateSeason(c *gin.Context) {

}

// swagger:route DELETE /seasons/{id} Season DeleteSeason
// Delete a season.
//
//	Parameters:
//		+ name: id
//		in: path
//		required: true
//		type: string
//
//	responses:
//		200:
//		500: ErrorResponse
func (h *SeasonHandler) deleteSeason(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		log.Printf("season id was not provided")
		c.JSON(http.StatusNotFound, models.NewErrorResponse(errors.New("missing id parameter")))
		return
	}

	err := h.Repo.DeleteSeason(id)
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	c.Status(http.StatusOK)
}
