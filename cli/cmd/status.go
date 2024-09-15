package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("syncClip is running...")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
