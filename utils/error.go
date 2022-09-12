package utils

import "fmt"

const RequestTimedOut = "request timed out"

func InvalidInput(input string) string {
	return fmt.Sprintf("Invalid Input %s", input)
}
