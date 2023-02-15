package main

import (
	"github.com/ivan-sabo/golang-playground/api/rest"
)

func main() {
	s := rest.NewRestServer()

	s.Run(":8080")
}
