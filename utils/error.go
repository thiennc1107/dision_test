package utils

import "fmt"

const RequestTimedOut = "request timed out"
const InfinityResult = "infinity result"
const NanResult = "NaN result"

func InvalidInput(input string) string {
	return fmt.Sprintf("Invalid Input %s", input)
}
