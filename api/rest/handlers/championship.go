package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan-sabo/golang-playground/api/rest/models"
	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
	"github.com/ivan-sabo/golang-playground/internal/championship/infrastructure/database/mysql/repository"
	"gorm.io/gorm"
)

type ChampionshipHandler struct {
	GinEngine *gin.Engine
	Repo      domain.ChampionshipRepo
}

func NewChampionshipHandler(ginEngine *gin.Engine, dbConn *gorm.DB) ChampionshipHandler {
	return ChampionshipHandler{
		GinEngine: ginEngine,
		Repo:      repository.NewChampionshipMySQLRepo(dbConn),
	}
}

func (gr *ChampionshipHandler) AddChampionshipRoutes() {
	c := gr.GinEngine.Group("/championships")

	c.GET("", gr.getChampionships)
	c.GET("/:id", gr.getChampionship)
	c.POST("", gr.postChampionship)
	c.PUT("/:id", gr.putChampionship)
	c.DELETE("/:id", gr.deleteChampionship)
}

// swagger:route GET /championships Championship getChampionships
// Get a list of Championships.
//
//	Produces:
//		- application/json
//
//	responses:
//		200: GetChampionshipsResponse
func (gr *ChampionshipHandler) getChampionships(c *gin.Context) {
	championships, err := gr.Repo.GetChampionships()
	if err != nil {
		log.Printf("An error occured: %v", err)
		return
	}

	c.JSON(http.StatusOK, models.NewGetChampionshipsResponse(championships))
}

// swagger:route GET /championship/{id} Championship getChampionship
// Get a single Championship by ID.
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
//		200: GetChampionshipResponse
//		404: ErrorResponse
//		500: ErrorResponse
func (ch *ChampionshipHandler) getChampionship(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		log.Printf("championship id was not provided")
		c.JSON(http.StatusNotFound, models.NewErrorResponse(errors.New("missing id parameter")))
		return
	}

	championship, err := ch.Repo.GetChampionship(id)
	if err == domain.ErrChampionshipNotFound {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusNotFound, models.NewErrorResponse(err))
		return
	}
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.NewGetChampionshipResponse(championship))
}

// swagger:route POST /championships Championship CreateChampionship
// Create a new Championship.
//
//	Consumes:
//		- application/json
//
//	Produces:
//		- application/json
//
//	responses:
//		200: PostChampionshipResponse
//		500: ErrorResponse
func (ch *ChampionshipHandler) postChampionship(c *gin.Context) {
	var cr models.PostChampionshipRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	championship, err := cr.ToEntity()
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	dc, err := ch.Repo.CreateChampionship(championship)
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, models.NewPostChampionshipResponse(dc))
}

// swagger:route PUT /championships/{id} Championship UpdateChampionship
// Update existing championship.
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
//		200: PutChampionshipResponse
//		404: ErrorResponse
//		500: ErrorResponse
func (ch *ChampionshipHandler) putChampionship(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		log.Printf("championship id was not provided")
		c.JSON(http.StatusNotFound, models.NewErrorResponse(errors.New("missing id parameter")))
		return
	}

	_, err := ch.Repo.GetChampionship(id)
	if err == domain.ErrChampionshipNotFound {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusNotFound, models.NewErrorResponse(err))
		return
	}
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	var championship models.PutChampionshipRequest
	err = c.ShouldBindJSON(&championship)
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	dc, err := ch.Repo.UpdateChampionship(id, championship.ToEntity())
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.NewPutChampionshipResponse(dc))
}

// swagger:route DELETE /championships/{id} Championship DeleteChampionship
// Delete a championship.
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
func (ch *ChampionshipHandler) deleteChampionship(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		log.Printf("championship id was not provided")
		c.JSON(http.StatusNotFound, models.NewErrorResponse(errors.New("missing id parameter")))
		return
	}

	err := ch.Repo.DeleteChampionship(id)
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	c.Status(http.StatusOK)
}
