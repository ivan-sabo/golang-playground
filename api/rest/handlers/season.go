package handlers

import (
	"github.com/gin-gonic/gin"
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

}

func (h *SeasonHandler) getSeason(c *gin.Context) {

}

func (h *SeasonHandler) createSeason(c *gin.Context) {

}

func (h *SeasonHandler) updateSeason(c *gin.Context) {

}

func (h *SeasonHandler) deleteSeason(c *gin.Context) {

}
