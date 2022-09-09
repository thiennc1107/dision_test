package server

import "fmt"

const InvalidController = "invalid controller version"

const NotImplemented = "controller not implemented"

func InvalidInput(input string) string {
	return fmt.Sprintf("Invalid Input %s", input)
}
