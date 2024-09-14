package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"syncClip/server/service"
	"syncClip/server/util"
)

// Register a new machine
func Register(c *gin.Context) {
	var req util.RegisterRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	_, err := service.GetOrAllocate(req.Host, req.Port, req.MAC)
	if err != nil {
		log.Printf("Failed to allocate ID for host %s: %v", req.Host, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	all := service.All()
	c.JSON(http.StatusOK, util.NewResponse(http.StatusOK, "success", all))
}
