package util

import (
	"net"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

func NewRestyClient(timeout time.Duration) *resty.Client {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   2 * time.Second, // Connection timeout
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   2 * time.Second, // TLS handshake timeout
		ResponseHeaderTimeout: 2 * time.Second, // Response header timeout
		ExpectContinueTimeout: 1 * time.Second,
	}
	client := resty.New()
	client.SetTransport(transport)
	client.SetTimeout(timeout)
	return client
}
