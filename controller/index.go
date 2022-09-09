package controller

import (
	"fmt"
	"worker/models"
	"worker/server"
)

type Controller interface {
	CalculateTest(a, b int16) (models.Data, error)
	log(msg string)
}

var enableLog bool

func ListController() {
	for k := range controllerMap {
		fmt.Println(k)
	}
}

func GetController(version string) (Controller, error) {
	var ok bool
	controller, ok := controllerMap[version]
	if !ok {
		return nil, fmt.Errorf(server.InvalidController)
	}
	return controller, nil
}

var controllerMap map[string]Controller

func RegisterController() {
	controllerMap = make(map[string]Controller)
	controllerMap["v1"] = &controllerV1{}
	controllerMap["v2"] = &controllerV2{}
}

func EnableLog() {
	enableLog = true
}
