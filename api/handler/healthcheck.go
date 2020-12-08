package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.SecureJSON(http.StatusOK, gin.H{
		"status": "running",
	})
}
