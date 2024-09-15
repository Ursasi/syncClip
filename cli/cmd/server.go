package cmd

import (
	"fmt"
	constant "syncClip"
	"syncClip/server"

	"github.com/spf13/cobra"
)

var serverConfig server.Config

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "manage server nodes",
}

var serverUpCmd = &cobra.Command{
	Use:   "up",
	Short: "up server node mode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("up server node mode...")
		server.StartServer(serverConfig)
	},
}
var serverDownCmd = &cobra.Command{
	Use:   "down",
	Short: "down server node mode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("down server node mode...")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.AddCommand(serverUpCmd)
	serverCmd.AddCommand(serverDownCmd)

	serverUpCmd.Flags().StringVarP(&serverConfig.Address, "address", "a", constant.DefaultListenAddress, "The HTTP Server listen address for peer service.")
	serverUpCmd.Flags().IntVarP(&serverConfig.Port, "port", "p", constant.DefaultServerListenPort, "The HTTP Server listen port for peer service.")
	serverUpCmd.Flags().IntVarP(&serverConfig.Concurrency, "max-concurrency", "c", constant.DefaultMaxConcurrency, fmt.Sprintf("The maximum number of concurrent connections the Server may serve, use the default value %d if <=0.", constant.DefaultMaxConcurrency))
	serverUpCmd.Flags().BoolVarP(&serverConfig.Logging, "api-logging", "l", true, "Enable API logging for peer request.")
	serverUpCmd.Flags().IntVarP(&serverConfig.CleanInterval, "clean-interval", "C", constant.DefaultCleanInterval, "The interval of cleaning the expired data in the bucket.")
}
