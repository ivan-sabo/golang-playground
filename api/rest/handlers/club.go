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

type ClubHandler struct {
	GinEngine *gin.Engine
	Repo      domain.ClubRepo
}

func NewClubHandler(ginEngine *gin.Engine, dbConn *gorm.DB) ClubHandler {
	return ClubHandler{
		GinEngine: ginEngine,
		Repo:      mysql.NewClubMySQLRepo(dbConn),
	}
}

func (ch *ClubHandler) AddClubRoutes() {
	c := ch.GinEngine.Group("/club")

	c.GET("/:id", ch.GetClub)
	c.GET("", ch.GetClubs)
}

// swagger:route GET /club Clubs getClubs
// Club holds all data relevant to a football club.
//
//	Produces:
//		- application/json
//
//	responses:
//		200: GetClubsResponse
func (ch *ClubHandler) GetClubs(c *gin.Context) {
	clubs, err := ch.Repo.GetClubs(domain.ClubFilter{})
	if err != nil {
		log.Printf("an error occured: %v", err)
		return
	}

	c.JSON(http.StatusOK, models.GetClubsResponse{
		Clubs: models.NewClubsDTO(clubs),
	})
}

// swagger:route GET /club/{id} Clubs getClub
// Club holds all data relevant to a football club.
//
//	Consumes:
//		- application/json
//
//	Produces:
//		- application/json
//
//	Parameters:
//		+ name: id
//		in: path
//		required: true
//		type: string
//
//	responses:
//		200: GetClubResponse
//		404: ErrorResponse
//		500: ErrorResponse
func (ch *ClubHandler) GetClub(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		log.Printf("club id was not provided")
		return
	}

	club, err := ch.Repo.GetClub(id)
	if err == domain.ErrClubNotFound {
		c.JSON(http.StatusNotFound, models.NewErrorResponse(err))
		return
	}
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.GetClubResponse{
		Club: models.NewClubDTO(club),
	})
}