package controller

import (
	"fmt"
	"worker/models"
	"worker/utils"
)

type Controller interface {
	CalculateTest(a, b int16) models.Data
}

var controllerConfig = &struct {
	enableLog bool
}{
	enableLog: false,
}

func EnableLog() {
	controllerConfig.enableLog = true
}

func log(message string) {
	if controllerConfig.enableLog {
		fmt.Println(message)
	}
}

func ListController() {
	for k := range controllerMap {
		println(k)
	}
}

func GetController(version string) (Controller, error) {
	controller, ok := controllerMap[version]
	if !ok {
		return nil, fmt.Errorf(utils.InvalidController)
	}
	return controller, nil
}

var controllerMap map[string]Controller

func RegisterController() {
	log("CTRL registered")
	controllerMap = make(map[string]Controller)
	controllerMap["v1"] = &controllerV1{}
}
