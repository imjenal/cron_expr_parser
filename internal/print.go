package internal

import (
	"fmt"
)

// PrintCronFields prints expanded cron fields
func PrintCronFields(fieldNames []string, fieldValues [][]int, command string) {
	fmt.Println("Expanded Cron Fields:")
	for i, fieldName := range fieldNames {
		if fieldName != "command" {
			fmt.Printf("%s: %v\n", fieldName, fieldValues[i])
		}
	}
	fmt.Printf("command: %s\n", command)
}
