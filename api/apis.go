package api

import (
	"github.com/gin-gonic/gin"
	"github.com/matheushr97/elastic-apm-with-golang/api/handler"
	"go.elastic.co/apm/module/apmgin"
)

func SetupRoutes(engine *gin.Engine) {
	apm := engine.Group("/apm")
	apm.Use(apmgin.Middleware(engine))
	apm.GET("/health-check", handler.HealthCheck)
	apm.GET("/call", handler.Call)
}
