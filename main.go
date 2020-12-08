package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matheushr97/elastic-apm-with-golang/api"
)

func main() {
	engine := gin.Default()
	api.SetupRoutes(engine)
	engine.Run()
}
