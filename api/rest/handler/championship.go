package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	models "github.com/ivan-sabo/golang-playground/api/rest/model"
	"github.com/ivan-sabo/golang-playground/internal"
	"github.com/ivan-sabo/golang-playground/internal/championship/application/service"
	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
	"gorm.io/gorm"
)

type ChampionshipHandler struct {
	ginEngine           *gin.Engine
	ChampionshipService domain.ChampionshipService
}

func NewChampionshipHandler(ginEngine *gin.Engine, dbConn *gorm.DB) ChampionshipHandler {
	return ChampionshipHandler{
		ginEngine:           ginEngine,
		ChampionshipService: service.NewChampionshipService(dbConn),
	}
}

func (gr *ChampionshipHandler) AddChampionshipRoutes() {
	c := gr.ginEngine.Group("/championships")
	{
		c.GET("", gr.getChampionships)
		c.GET("/:championshipID", gr.getChampionship)
		c.POST("", gr.postChampionship)
		c.PUT("/:championshipID", gr.putChampionship)
		c.DELETE("/:championshipID", gr.deleteChampionship)
	}

	c = gr.ginEngine.Group("/championships/:championshipID")
	{
		c.GET("/seasons", gr.getChampionshipsSeasons)
		c.POST("/seasons/:seasonID", gr.registerSeason)
	}
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
	// @todo: implement filter
	championships, err := gr.ChampionshipService.GetChampionships(domain.ChampionshipFilter{})
	if err != nil {
		log.Printf("An error occured: %v", err)
		return
	}

	c.JSON(http.StatusOK, models.NewGetChampionshipsResponse(championships))
}

// swagger:route GET /championships/{id} Championship getChampionship
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
//		400: ErrorResponse
//		404: ErrorResponse
//		500: ErrorResponse
func (ch *ChampionshipHandler) getChampionship(c *gin.Context) {
	id, exist := c.Params.Get("championshipID")
	if !exist {
		log.Printf("championship id was not provided")
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(errors.New("missing id parameter")))
		return
	}

	championship, err := ch.ChampionshipService.GetChampionship(id)
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

	dc, err := ch.ChampionshipService.CreateChampionship(championship)
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
//		400: ErrorResponse
//		404: ErrorResponse
//		500: ErrorResponse
func (ch *ChampionshipHandler) putChampionship(c *gin.Context) {
	id, exist := c.Params.Get("championshipID")
	if !exist {
		log.Printf("championship id was not provided")
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(errors.New("missing id parameter")))
		return
	}

	var championship models.PutChampionshipRequest
	err := c.ShouldBindJSON(&championship)
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	dc, err := ch.ChampionshipService.UpdateChampionship(id, championship.ToEntity())
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
//		400: ErrorResponse
//		500: ErrorResponse
func (ch *ChampionshipHandler) deleteChampionship(c *gin.Context) {
	id, exist := c.Params.Get("championshipID")
	if !exist {
		log.Printf("championship id was not provided")
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(errors.New("missing id parameter")))
		return
	}

	err := ch.ChampionshipService.DeleteChampionship(id)
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	c.Status(http.StatusOK)
}

// swagger:route POST /championships/{championshipID}/seasons/{seasonID} Championship RegisterSeason
// Register Season for Championship.
//
//	Parameters:
//		+ name: championshipID
//		in: path
//		required: true
//		type: string
//
//		+ name: seasonID
//		in: path
//		required: true
//		type: string
//
//	Produces:
//		- application/json
//
//	responses:
//		200: RegisterSeasonResponse
//		400: ErrorResponse
//		500: ErrorResponse
func (ch *ChampionshipHandler) registerSeason(c *gin.Context) {
	championshipID, exist := c.Params.Get("championshipID")
	if !exist {
		log.Printf("championship id was not provided")
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(errors.New("missing championshipID parameter")))
		return
	}

	seasonID, exist := c.Params.Get("seasonID")
	if !exist {
		log.Printf("season id was not provided")
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(errors.New("missing seasonID parameter")))
		return
	}

	cs, err := ch.ChampionshipService.RegisterSeason(championshipID, seasonID)
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, models.NewChampionshipSeason(cs))
}

// swagger:route GET /championships/{championshipID}/seasons Championship getChampionshipsSeasons
// Get filtered ChampionshipsSeasons.
//
//	Produces:
//		- application/json
//
//	Parameters:
//		+ name: championshipID
//		in: path
//		required: true
//		type: string
//
//		+ name: seasonID
//		in: query
//		required: false
//		type: string
//
//	responses:
//		200: GetChampionshipsSeasonsResponse
//		400: ErrorResponse
//		500: ErrorResponse
func (ch *ChampionshipHandler) getChampionshipsSeasons(c *gin.Context) {
	csFilter := domain.ChampionshipSeasonFilter{}

	championshipID, exist := c.Params.Get("championshipID")
	if exist {
		championshipUUID, err := uuid.Parse(championshipID)
		if err != nil {
			log.Printf("an error occured: %v", err)
			c.JSON(http.StatusBadRequest, models.NewErrorResponse(internal.ErrInvalidUUID))
			return
		}

		csFilter.ChampionshipID = championshipUUID.String()
	}

	seasonID := c.Query("seasonID")
	if seasonID != "" {
		seasonUUID, err := uuid.Parse(seasonID)
		if err != nil {
			log.Printf("an error occured: %v", err)
			c.JSON(http.StatusBadRequest, models.NewErrorResponse(internal.ErrInvalidUUID))
			return
		}

		csFilter.SeasonID = seasonUUID.String()
	}

	csss, err := ch.ChampionshipService.GetChampionshipsSeasons(csFilter)
	if err != nil {
		log.Printf("an error occured: %v", err)
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.NewChampionshipsSeasons(csss))
}
