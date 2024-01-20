package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{Use: "cron_expr_parser"}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// Handle error
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
