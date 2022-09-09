package controller

import (
	"fmt"
	"math"
	"worker/models"
)

type controllerV1 struct {
}

func (c *controllerV1) log(msg string) {
	if enableLog {
		fmt.Println("log test")
	}
}

func (c *controllerV1) CalculateTest(a, b int16) (models.Data, error) {
	c.log("log test")
	return models.Data{
		F1: int64(a + b),
		F2: int64(a * b),
		F3: math.Exp(float64(a)) * math.Exp(float64(b)),
		F4: math.Exp(float64(a)) * (-math.Exp(float64(b))),
	}, nil
}
