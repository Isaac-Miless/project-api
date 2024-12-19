package main

import (
	"log"
	"net/http"
	"os"
	"project-api/database"
	"project-api/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file (only locally)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	database.Connect()
	router := routes.Initialize()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server started at :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
