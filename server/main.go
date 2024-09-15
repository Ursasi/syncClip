package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/automaxprocs/maxprocs"
	"syncClip/server/handler"
	"syncClip/server/service"
)

type ServerConfig struct {
	Address     string
	Port        int
	Concurrency int
	Logging     bool
}

const (
	defaultPort           = 27281
	defaultMaxConcurrency = 8
)

var serverConfig ServerConfig

func init() {
	pflag.StringVar(&serverConfig.Address, "address", "0.0.0.0", "The HTTP Server listen address for syncClip service.")
	pflag.IntVar(&serverConfig.Port, "port", defaultPort, "The HTTP Server listen port for kb-agent service.")
	pflag.IntVar(&serverConfig.Concurrency, "max-concurrency", defaultMaxConcurrency,
		fmt.Sprintf("The maximum number of concurrent connections the Server may serve, use the default value %d if <=0.", defaultMaxConcurrency))
	pflag.BoolVar(&serverConfig.Logging, "api-logging", true, "Enable API logging for kb-agent request.")
}

func main() {
	_, _ = maxprocs.Set()

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	pflag.Parse()

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(errors.Wrap(err, "fatal error viper bindPFlags"))
	}

	fmt.Printf("Server is starting with config: %+v\n", serverConfig)

	startServer()
}

func startServer() {
	r := gin.Default()

	r.POST("/register", handler.Register)
	r.POST("/get", handler.Get)
	r.POST("/all", handler.All)
	r.POST("/connect", handler.Connect)
	r.POST("/disconnect", handler.Disconnect)

	service.InitBucket()
	addr := fmt.Sprintf("%s:%d", serverConfig.Address, serverConfig.Port)
	if err := r.Run(addr); err != nil {
		panic(errors.Wrap(err, "failed to start the server"))
	}
}
