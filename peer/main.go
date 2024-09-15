package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/automaxprocs/maxprocs"
	"runtime"
	"syncClip/peer/handler"
)

type PeerConfig struct {
	Address       string
	ServerAddress string
	Port          int
	Concurrency   int
	Logging       bool
}

const (
	defaultPort           = 27281
	defaultMaxConcurrency = 8
)

var peerConfig PeerConfig
var clipChan = make(chan string, 2048)

func init() {
	pflag.StringVar(&peerConfig.Address, "address", "0.0.0.0", "The HTTP Server listen address for peer.")
	pflag.StringVar(&peerConfig.Address, "server-address", "0.0.0.0", "The HTTP Server listen address for syncClip service.")
	pflag.IntVar(&peerConfig.Port, "port", defaultPort, "The HTTP Server listen port for kb-agent service.")
	pflag.IntVar(&peerConfig.Concurrency, "max-concurrency", defaultMaxConcurrency,
		fmt.Sprintf("The maximum number of concurrent connections the Server may serve, use the default value %d if <=0.", defaultMaxConcurrency))
	pflag.BoolVar(&peerConfig.Logging, "api-logging", true, "Enable API logging for kb-agent request.")
}

func main() {
	_, _ = maxprocs.Set()

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	pflag.Parse()

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(errors.Wrap(err, "fatal error viper bindPFlags"))
	}

	fmt.Printf("Server is starting with config: %+v\n", peerConfig)
	startServer()
}

func startServer() {
	r := gin.Default()

	r.POST("/receive", handler.Receive)

	WatchClip()
	addr := fmt.Sprintf("%s:%d", peerConfig.Address, peerConfig.Port)
	if err := r.Run(addr); err != nil {
		panic(errors.Wrap(err, "failed to start the server"))
	}
}

func WatchClip() {
	clipChan = make(chan string)
	go func() {
		if runtime.GOOS == "darwin" {
			WatchClipboard()
		}
	}()
	go func() {
		for {
			select {
			case msg := <-clipChan:
				if msg != "" {
					// TODO need to send to server
					fmt.Println("send to server:", msg)
				}
			}
		}
	}()
}
