package handler

import (
	"net/http"
	"syncClip/server/service"
	"syncClip/util"

	"github.com/gin-gonic/gin"
)

// Get by ID and return the machine info
func Get(c *gin.Context) {
	var req util.GetRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	board := service.Get(req.IP, req.Port)
	c.JSON(http.StatusOK, util.NewResponse(http.StatusOK, "success", board))
}

// All returns all machine info in LAN
func All(c *gin.Context) {
	allInLan := service.All()
	c.JSON(http.StatusOK, util.NewResponse(http.StatusOK, "success", allInLan))
}

func Probe(c *gin.Context) {
	var req util.ProbeRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	boards := service.Probe(req.ID)
	c.JSON(http.StatusOK, util.NewResponse(http.StatusOK, "success", boards))
}
