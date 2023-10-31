package main

import (
	"log"

	"github.com/SicParv1sMagna/HappyPetsBackend/internal/app"
)

// @title HappyPets RestAPI
// @version 1.0
// @description API server for Native HappyPets application

// @host http://localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	log.Println("API start!")

	application, err := app.New()

	if err != nil {
		log.Fatal(err)
	}

	application.StartServer()

	log.Println("API is shitting down!")
}
