package controller

import (
	"fmt"
	"worker/models"
	"worker/server"
)

type controllerV2 struct {
}

func (c *controllerV2) CalculateTest(a, b int16) (models.Data, error) {
	return models.Data{}, fmt.Errorf(server.NotImplemented)
}

func (c *controllerV2) log(msg string) {
}
