package main

import (
	"log"

	"github.com/glebb1331/middleware/internal/pkg/app"
)

func main() {

	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
