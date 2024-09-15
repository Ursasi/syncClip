package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Connect to a machine
func Connect(c *gin.Context) {
	data := map[string]string{
		"key":    "value",
		"status": "success",
	}
	c.JSON(http.StatusOK, data)
}

// Disconnect from a machine
func Disconnect(c *gin.Context) {
	data := map[string]string{
		"key":    "value",
		"status": "success",
	}
	c.JSON(http.StatusOK, data)
}
