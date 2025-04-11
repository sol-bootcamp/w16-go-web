package main

import (
	"bootcamp-web/cmd/server"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// r := chi.NewRouter()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	API_TOKEN := os.Getenv("API_TOKEN")

	cfg := &server.ConfigServerChi{
		ServerAddress:  ":" + PORT,
		LoaderFilePath: "products.json",
	}

	log.Println("Starting server on :" + PORT)
	app := server.NewServerChi(cfg)
	// - run
	if err := app.Run(API_TOKEN); err != nil {
		fmt.Println(err)
		return
	}

}
