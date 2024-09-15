package cmd

import (
	"fmt"
	constant "syncClip"
	"syncClip/peer"
	"syncClip/util"
	"time"

	"github.com/spf13/cobra"
)

var peerConfig peer.Config

var peersWatch bool

var peerCmd = &cobra.Command{
	Use:   "peer",
	Short: "manage peer nodes",
}
var peersCmd = &cobra.Command{
	Use:   "peers",
	Short: "list connected peer nodes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("all peer nodes:")
		table := util.ConfigTable()
		data := [][]string{
			{"10.233.4.1", "1234", "SKFNCKFLS.SKGNCJH"},
			{"10.233.4.2", "4502", "YXNVKNKFK.SXKFNXJ"},
			{"10.233.4.3", "3501", "SNCHSKLFM.LKFHXXJ"},
			{"10.233.4.4", "5042", "QFCKVMFJV.ONFMSKL"},
		}
		table.AppendBulk(data)
		for {
			util.ClearScreen()
			util.ShowPeers(table)
			time.Sleep(3 * time.Second)
		}
	},
}

var peerUpCmd = &cobra.Command{
	Use:   "up",
	Short: "start peer node mode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("peer starting...")
		peer.StartPeer(peerConfig)
	},
}
var peerDownCmd = &cobra.Command{
	Use:   "down",
	Short: "close peer node mode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("peer closing...")
	},
}

var peerShowCmd = &cobra.Command{
	Use:   "show",
	Short: "show peer node info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("peer node info")
	},
}

func init() {
	rootCmd.AddCommand(peerCmd)
	rootCmd.AddCommand(peersCmd)
	peerCmd.AddCommand(peerUpCmd)
	peerCmd.AddCommand(peerDownCmd)
	peerCmd.AddCommand(peerShowCmd)

	peerUpCmd.Flags().StringVarP(&peerConfig.Address, "address", "a", constant.DefaultListenAddress, "The HTTP Server listen address for peer service.")
	peerUpCmd.Flags().IntVarP(&peerConfig.Port, "port", "p", constant.DefaultPeerListenPort, "The HTTP Server listen port for peer service.")
	peerUpCmd.Flags().StringVarP(&peerConfig.ServerAddress, "server-address", "s", constant.DefaultServerAddress, "The HTTP Server listen address for syncClip service.")
	peerUpCmd.Flags().IntVarP(&peerConfig.ServerPort, "server-port", "P", constant.DefaultServerListenPort, "The HTTP Server listen port for syncClip service.")
	peerUpCmd.Flags().IntVarP(&peerConfig.Concurrency, "max-concurrency", "c", constant.DefaultMaxConcurrency, fmt.Sprintf("The maximum number of concurrent connections the Server may serve, use the default value %d if <=0.", constant.DefaultMaxConcurrency))
	peerUpCmd.Flags().BoolVarP(&peerConfig.Logging, "api-logging", "l", true, "Enable API logging for peer request.")

	peersCmd.Flags().BoolVarP(&peersWatch, "watch", "w", true, "Continuous monitoring.")
}
