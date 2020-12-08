package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/matheushr97/elastic-apm-with-golang/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	engine := gin.Default()
	api.SetupRoutes(engine)
	engine.Run()
}
