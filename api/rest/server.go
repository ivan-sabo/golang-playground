package rest

import (
	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RestServer struct {
	Router   *gin.Engine
	SqlStore *gorm.DB
}

func NewRestServer(dbConn *gorm.DB) *RestServer {
	s := gin.Default()
	s.Use(cors.Default())

	attachRoutes(s, dbConn)

	return &RestServer{
		Router:   s,
		SqlStore: dbConn,
	}
}
