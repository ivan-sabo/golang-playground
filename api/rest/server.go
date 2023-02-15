package rest

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func NewRestServer() *gin.Engine {
	s := gin.Default()

	s.Use(cors.Default())

	attachRoutes(s)

	return s
}
