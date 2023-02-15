package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan-sabo/golang-playground/api/rest/models"
)

// swagger:route GET /championship Championship getChampionship
// Championship holds all data relevant to a soccer league.
//
//	Produces:
//		- application/json
//
//	responses:
//		200: ChampionshipResponse
func GetChampionship(c *gin.Context) {
	ch := models.ChampionshipResponseWrapper{
		Body: models.Championship{
			Clubs: models.Clubs{
				{Name: "NK Radnicki"},
			},
		},
	}
	c.JSON(http.StatusOK, ch)
}
