package server

import (
	"flag"
	"fmt"
	"syncClip/server/handler"
	"syncClip/server/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/automaxprocs/maxprocs"
)

type Config struct {
	Address       string
	Port          int
	Concurrency   int
	Logging       bool
	CleanInterval int
}

var serverConfig Config

func StartServer(config Config) {
	_, _ = maxprocs.Set()
	serverConfig = config

	initFlags()
	go startCleaner()
	startServer()
}

func initFlags() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(errors.Wrap(err, "fatal error viper bindPFlags"))
	}
}
func startCleaner() {
	ticker := time.NewTicker(time.Duration(serverConfig.CleanInterval))
	defer ticker.Stop()

	for {
		<-ticker.C
		ticker.Stop()
		service.B.Clean()
		ticker = time.NewTicker(time.Duration(serverConfig.CleanInterval))
	}
}
func startServer() {
	r := gin.Default()

	r.POST("/register", handler.Register)
	r.POST("/get", handler.Get)
	r.POST("/all", handler.All)
	r.POST("/connect", handler.Connect)
	r.POST("/disconnect", handler.Disconnect)
	r.POST("/probe", handler.Probe)

	service.InitBucket()
	addr := fmt.Sprintf("%s:%d", serverConfig.Address, serverConfig.Port)
	if err := r.Run(addr); err != nil {
		panic(errors.Wrap(err, "failed to start the server"))
	}
}
