package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"syncClip/peer/service"
	"syncClip/util"
)

func Receive(c *gin.Context) {
	var req util.ReceiveRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	flag, err := service.Receive(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, util.NewResponse(http.StatusOK, "success", flag))
}
