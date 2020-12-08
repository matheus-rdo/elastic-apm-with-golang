package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Call(c *gin.Context) {
	time.Sleep(1 * time.Second)
	resultParam := c.Query("result")
	switch resultParam {
	case "panic":
		panic("this is a panic!!")
	case "error":
		c.AbortWithError(http.StatusInternalServerError, errors.New("this is an fatal error"))
	default:
		c.SecureJSON(http.StatusOK, gin.H{
			"status": "running",
		})
	}
}
