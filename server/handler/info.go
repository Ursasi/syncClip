package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"syncClip/server/service"
	"syncClip/util"
)

// Get by ID and return the machine info
func Get(c *gin.Context) {
	var req util.GetRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	board := service.Get(req.IP, req.Port, req.MAC)
	c.JSON(http.StatusOK, util.NewResponse(http.StatusOK, "success", board))
}

// All returns all machine info in LAN
func All(c *gin.Context) {
	allInLan := service.All()
	c.JSON(http.StatusOK, util.NewResponse(http.StatusOK, "success", allInLan))
}
