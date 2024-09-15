package util

import (
	"net"

	"github.com/gin-gonic/gin"
)

type GetRequest struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}
type ConnectRequest struct {
	SID string `json:"sid"`
	DID string `json:"did"`
}

type DisconnectRequest struct {
	SID string `json:"sid"`
	DID string `json:"did"`
}

type RegisterRequest struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
type ReceiveRequest struct {
	Msg string
}

type ProbeRequest struct {
	ID string `json:"id"`
}

func ParseIPNPort(c *gin.Context) (string, string, error) {
	ip := c.ClientIP()
	clientAddr := c.Request.RemoteAddr
	ip, port, err := net.SplitHostPort(clientAddr)
	if err != nil {
		return "", "", err
	}
	return ip, port, nil
}
