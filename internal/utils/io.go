package utils

import "fmt"

func RequestConfirmation(message string) bool {
	var confirmation string

	fmt.Printf("%s", message)
	fmt.Scanf("%s", &confirmation)

	if confirmation == "yes" || confirmation == "y" {
		return true
	}

	return false
}
