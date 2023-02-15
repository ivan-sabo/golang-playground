package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/ivan-sabo/golang-playground/api/rest/handlers"
)

func attachRoutes(g *gin.Engine) {
	c := g.Group("/championship")

	c.GET("", handlers.GetChampionship)
}
