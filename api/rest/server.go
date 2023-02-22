package rest

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"gorm.io/gorm"
)

func NewRestServer(dbConn *gorm.DB) *gin.Engine {
	s := gin.Default()

	s.Use(cors.Default())

	attachRoutes(s, dbConn)

	return s
}
