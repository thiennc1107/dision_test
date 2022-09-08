package controller

import (
	"math"
	"worker/models"
)

type controllerV1 struct {
}

func (c *controllerV1) CalculateTest(a, b int16) models.Data {
	return models.Data{
		F1: a + b,
		F2: a * b,
		F3: math.Exp(float64(a)) * math.Exp(float64(b)),
		F4: math.Exp(float64(a)) * (-math.Exp(float64(b))),
	}
}
