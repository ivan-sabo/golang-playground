package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan-sabo/golang-playground/api/rest/models"
)

// swagger:route GET /championship championship-tag
// Championship holds all data relevant to a soccer league.
// responses:
//
//	200: championshipResponse
func GetChampionship(c *gin.Context) {
	ch := models.ChampionshipResponse{
		Clubs: models.Clubs{
			{Name: "NK Radnicki"},
		},
	}
	c.JSON(http.StatusOK, ch)
}
