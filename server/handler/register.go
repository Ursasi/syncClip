package handler

import (
	"log"
	"net/http"
	"syncClip/server/service"
	"syncClip/util"

	"github.com/gin-gonic/gin"
)

// Register a new machine
func Register(c *gin.Context) {
	ip, port, err := util.ParseIPNPort(c)
	if err != nil {
		log.Printf("Failed to parse IP and port: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	ID, err := service.GetOrAllocate(ip, port)
	if err != nil {
		log.Printf("Failed to allocate ID for host %s: %v", ip, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	boards := service.All()
	c.JSON(http.StatusOK, util.NewResponse(http.StatusOK, "success", util.RegisterResponse{
		ID:     ID,
		Boards: boards,
	}))
}
