package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/ivan-sabo/golang-playground/api/rest/model"
	"github.com/ivan-sabo/golang-playground/internal/championship/application/service"
	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
	"gorm.io/gorm"
)

type SeasonHandler struct {
	GinEngine     *gin.Engine
	seasonService domain.SeasonService
}

func NewSeasonHandler(g *gin.Engine, conn *gorm.DB) SeasonHandler {
	return SeasonHandler{
		GinEngine:     g,
		seasonService: service.NewSeasonService(conn),
	}
}

func (h *SeasonHandler) AddSeasonRoutes() {
	g := h.GinEngine.Group("/seasons")

	g.GET("", h.getSeasons)
	g.GET("/:id", h.getSeason)
	g.POST("", h.createSeason)
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
func (h *SeasonHandler) getSeasons(ctx *gin.Context) {
	seasons, err := h.seasonService.GetSeasons(ctx, domain.SeasonFilter{})
	if err != nil {
		log.Printf("an error occured: %v", err)
		ctx.JSON(http.StatusNotFound, models.NewErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.NewGetSeasonsResponse(seasons))
}

// swagger:route GET /seasons/{id} Season getSeason
// Get a single Season by ID.
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
//		200: GetSeasonResponse
//		400: ErrorResponse
//		404: ErrorResponse
//		500: ErrorResponse
func (h *SeasonHandler) getSeason(ctx *gin.Context) {
	id, exist := ctx.Params.Get("id")
	if !exist {
		log.Printf("season id was not provided")
		ctx.JSON(http.StatusBadRequest, models.NewErrorResponse(errors.New("missing id parameter")))
		return
	}
	season, err := h.seasonService.GetSeason(ctx, id)
	if err == domain.ErrSeasonNotFound {
		log.Printf("an error occured: %v", err)
		ctx.JSON(http.StatusNotFound, models.NewErrorResponse(err))
		return
	}
	if err != nil {
		log.Printf("an error occured: %v", err)
		ctx.JSON(http.StatusNotFound, models.NewErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.NewGetSeasonResponse(season))
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
func (h *SeasonHandler) createSeason(ctx *gin.Context) {
	var sr models.PostSeasonRequest
	err := ctx.ShouldBindJSON(&sr)
	if err != nil {
		log.Printf("an error occured: %v", err)
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	season, err := sr.ToEntity()
	if err != nil {
		log.Printf("an error occured: %v", err)
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	ds, err := h.seasonService.CreateSeason(ctx, season)
	if err != nil {
		log.Printf("an error occured: %v", err)
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, models.NewPostSeasonResponse(ds))
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
func (h *SeasonHandler) deleteSeason(ctx *gin.Context) {
	id, exist := ctx.Params.Get("id")
	if !exist {
		log.Printf("season id was not provided")
		ctx.JSON(http.StatusNotFound, models.NewErrorResponse(errors.New("missing id parameter")))
		return
	}

	err := h.seasonService.DeleteSeason(ctx, id)
	if err != nil {
		log.Printf("an error occured: %v", err)
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}
