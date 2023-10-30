package main

import (
	"log"

	"github.com/SicParv1sMagna/HappyPetsBackend/internal/app"
)

func main() {
	log.Println("API start!")

	application, err := app.New()

	if err != nil {
		log.Fatal(err)
	}

	application.StartServer()

	log.Println("API is shitting down!")
}
