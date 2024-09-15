package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "syncClip",
	Short:   "a tool to sync clipboard content between devices",
	Long:    `syncClip is a tool to sync clipboard content between devices. It can be used to sync text`,
	Aliases: []string{"synclip", "sp"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("syncClip is running...")
	},
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
