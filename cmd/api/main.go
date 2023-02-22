package main

import (
	"github.com/ivan-sabo/golang-playground/api/rest"
	"github.com/ivan-sabo/golang-playground/internal/championship/infrastructure/database/mysql"
)

func main() {
	dbConf := mysql.Config{
		DbName:   "golang_playground",
		Username: "root",
		Password: "test",
		Host:     "localhost",
	}

	dbConn, err := mysql.GetConnection(dbConf)
	if err != nil {
		panic(err)
	}

	s := rest.NewRestServer(dbConn)

	s.Router.Run(":8080")
}
