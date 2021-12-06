package main

import (
	"log"

	"imdb/config"

	"imdb/model"
	"imdb/router"

	"github.com/joho/godotenv"
)

func main() {
	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// connect db
	db := config.Connect()
	db.AutoMigrate(&model.Log{})

	r := router.SetupRouter()

	// running
	r.Run(":12006")

}
