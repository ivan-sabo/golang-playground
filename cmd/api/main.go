package main

import (
	"github.com/ivan-sabo/golang-playground/api/rest"
	"github.com/ivan-sabo/golang-playground/internal/championship/infrastructure/database/mysql"
)

func main() {
	dbConf := mysql.Config{
		DbName:   "golang-playground",
		Username: "root",
		Password: "test",
		Host:     "localhost",
	}

	dbConn, err := mysql.GetConnection(dbConf)
	if err != nil {
		panic(err)
	}

	s := rest.NewRestServer(dbConn)

	s.Run(":8080")
}
