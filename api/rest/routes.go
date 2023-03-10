package rest

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/ivan-sabo/golang-playground/api/rest/handler"
	"gorm.io/gorm"
)

func attachRoutes(g *gin.Engine, dbConn *gorm.DB) {
	championshipRouter := handlers.NewChampionshipHandler(g, dbConn)
	championshipRouter.AddChampionshipRoutes()

	clubRouter := handlers.NewClubHandler(g, dbConn)
	clubRouter.AddClubRoutes()

	seasonRouter := handlers.NewSeasonHandler(g, dbConn)
	seasonRouter.AddSeasonRoutes()
}
