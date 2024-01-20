package cmd

import (
	"cron_expr_parser/internal"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(expandCmd)
}

var expandCmd = &cobra.Command{
	Use:   "expand",
	Short: "Expand a cron expression",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cronExpression := args[0]

		if err := internal.ValidateCronExpression(cronExpression); err != nil {
			fmt.Println("Invalid cron expression:", err)
			return
		}

		fieldNames, fieldValues, command, err := internal.ExpandCronExpression(cronExpression)
		if err != nil {
			fmt.Println("Error expanding cron expression:", err)
			return
		}
		internal.PrintCronFields(fieldNames, fieldValues, command)
	},
}
