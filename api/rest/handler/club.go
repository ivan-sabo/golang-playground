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

type ClubHandler struct {
	ginEngine   *gin.Engine
	clubService domain.ClubService
}

func NewClubHandler(ginEngine *gin.Engine, dbConn *gorm.DB) ClubHandler {
	return ClubHandler{
		ginEngine:   ginEngine,
		clubService: service.NewClubService(dbConn),
	}
}

func (ch *ClubHandler) AddClubRoutes() {
	c := ch.ginEngine.Group("/clubs")

	c.GET("/:id", ch.getClub)
	c.PUT("/:id", ch.putClub)
	c.GET("", ch.getClubs)
	c.POST("", ch.postClub)
	c.DELETE("/:id", ch.deleteClub)
}

// swagger:route GET /clubs Club getClubs
// Get a list of clubs.
//
//	Produces:
//		- application/json
//
//	responses:
//		200: GetClubsResponse
func (ch *ClubHandler) getClubs(ctx *gin.Context) {
	clubs, err := ch.clubService.GetClubs(ctx, domain.ClubFilter{})
	if err != nil {
		log.Printf("an error occured: %v", err)
		ctx.JSON(http.StatusNotFound, models.NewErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.NewGetClubsResponse(clubs))
}

// swagger:route GET /clubs/{id} Club getClub
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
//		400: ErrorResponse
//		404: ErrorResponse
//		500: ErrorResponse
func (ch *ClubHandler) getClub(ctx *gin.Context) {
	id, exist := ctx.Params.Get("id")
	if !exist {
		log.Printf("club id was not provided")
		ctx.JSON(http.StatusBadRequest, models.NewErrorResponse(errors.New("missing id parameter")))
		return
	}

	club, err := ch.clubService.GetClub(ctx, id)
	if err == domain.ErrClubNotFound {
		log.Printf("an error occured: %v", err)
		ctx.JSON(http.StatusNotFound, models.NewErrorResponse(err))
		return
	}
	if err != nil {
		log.Printf("an error occured: %v", err)
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.GetClubResponse{
		Club: models.NewClubDTO(club),
	})
}

// swagger:route POST /clubs Club CreateClub
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
func (ch *ClubHandler) postClub(ctx *gin.Context) {
	var cr models.PostClubRequest
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		log.Printf("an error occured: %v", err)
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	club, err := cr.ToEntity()
	if err != nil {
		log.Printf("an error occured: %v", err)
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	dc, err := ch.clubService.CreateClub(ctx, club)
	if err != nil {
		log.Printf("an error occured: %v", err)
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, models.NewPostClubResponse(dc))
}

// swagger:route PUT /clubs/{id} Club UpdateClub
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
//		400: ErrorResponse
//		404: ErrorResponse
//		500: ErrorResponse
func (ch *ClubHandler) putClub(ctx *gin.Context) {
	id, exist := ctx.Params.Get("id")
	if !exist {
		log.Printf("club id was not provided")
		ctx.JSON(http.StatusBadRequest, models.NewErrorResponse(errors.New("missing id parameter")))
		return
	}

	_, err := ch.clubService.GetClub(ctx, id)
	if err == domain.ErrClubNotFound {
		log.Printf("an error occured: %v", err)
		ctx.JSON(http.StatusNotFound, models.NewErrorResponse(err))
		return
	}
	if err != nil {
		log.Printf("an error occured: %v", err)
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	var club models.PutClubRequest
	err = ctx.ShouldBindJSON(&club)
	if err != nil {
		log.Printf("an error occured: %v", err)
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	dc, err := ch.clubService.UpdateClub(ctx, id, club.ToEntity())
	if err != nil {
		log.Printf("an error occured: %v", err)
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.NewPutClubResponse(dc))
}

// swagger:route DELETE /clubs/{id} Club DeleteClub
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
//		400: ErrorResponse
//		500: ErrorResponse
func (ch *ClubHandler) deleteClub(ctx *gin.Context) {
	id, exist := ctx.Params.Get("id")
	if !exist {
		log.Printf("club id was not provided")
		ctx.JSON(http.StatusBadRequest, models.NewErrorResponse(errors.New("missing id parameter")))
		return
	}

	err := ch.clubService.DeleteClub(ctx, id)
	if err != nil {
		log.Printf("an error occured: %v", err)
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}
