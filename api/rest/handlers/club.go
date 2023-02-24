package handlers

import (
	"errors"
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

	c.GET("/:id", ch.getClub)
	c.PUT("/:id", ch.putClub)
	c.GET("", ch.getClubs)
	c.POST("", ch.postClub)
	c.DELETE("/:id", ch.deleteClub)
}

// swagger:route GET /club Clubs getClubs
// Get a list of clubs.
//
//	Produces:
//		- application/json
//
//	responses:
//		200: GetClubsResponse
func (ch *ClubHandler) getClubs(c *gin.Context) {
	clubs, err := ch.Repo.GetClubs(domain.ClubFilter{})
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusNotFound, models.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.GetClubsResponse{
		Clubs: models.NewClubsDTO(clubs),
	})
}

// swagger:route GET /club/{id} Clubs getClub
// Get a single club by ID.
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
func (ch *ClubHandler) getClub(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		log.Printf("club id was not provided")
		c.JSON(http.StatusNotFound, models.NewErrorResponse(errors.New("missing id parameter")))
		return
	}

	club, err := ch.Repo.GetClub(id)
	if err == domain.ErrClubNotFound {
		log.Printf("an error occured: %v", err)
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

// swagger:route POST /club Clubs CreateClub
// Create a new Club.
//
//	Consumes:
//		- application/json
//
//	Produces:
//		- application/json
//
//	responses:
//		200: PostClubResponse
//		500: ErrorResponse
func (ch *ClubHandler) postClub(c *gin.Context) {
	var club models.PostClubRequest
	err := c.ShouldBindJSON(&club)
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	dc, err := ch.Repo.CreateClub(club.ToEntity())
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, models.GetClubResponse{
		Club: models.NewClubDTO(dc),
	})
}

// swagger:route PUT /club/{id} Clubs UpdateClub
// Update existing club.
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
//		200: PutClubResponse
//		404: ErrorResponse
//		500: ErrorResponse
func (ch *ClubHandler) putClub(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		log.Printf("club id was not provided")
		c.JSON(http.StatusNotFound, models.NewErrorResponse(errors.New("missing id parameter")))
		return
	}

	_, err := ch.Repo.GetClub(id)
	if err == domain.ErrClubNotFound {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusNotFound, models.NewErrorResponse(err))
		return
	}
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	var club models.PutClubRequest
	err = c.ShouldBindJSON(&club)
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	dc, err := ch.Repo.UpdateClub(id, club.ToEntity())
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.NewPutClubResponse(dc))
}

// swagger:route DELETE /club/{id} Clubs DeleteClub
// Delete a club.
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
func (ch *ClubHandler) deleteClub(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		log.Printf("club id was not provided")
		c.JSON(http.StatusNotFound, models.NewErrorResponse(errors.New("missing id parameter")))
		return
	}

	err := ch.Repo.DeleteClub(id)
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	c.Status(http.StatusOK)
}
