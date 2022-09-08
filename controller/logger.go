package controller

import "fmt"

func EnableLog() {
	controllerConfig.enableLog = true
}

func log(message string) {
	if controllerConfig.enableLog {
		fmt.Println(message)
	}
}
