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
//		200: GetChampionshipResponse
func GetChampionship(c *gin.Context) {
	ch := models.GetChampionshipResponse{
		Championship: models.ChampionshipDTO{
			Clubs: models.ClubsDTO{
				{Name: "NK Radnicki"},
			},
			Name: "UEFA",
		},
	}
	c.JSON(http.StatusOK, ch)
}
