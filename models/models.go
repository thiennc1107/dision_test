package models

import (
	"fmt"
	"math"
)

type Data struct {
	F1 int64   `json:"f1"`
	F2 int64   `json:"f2"`
	F3 float64 `json:"f3"`
	F4 float64 `json:"f4"`
}

func (d Data) CheckInvalidResult() (Data, error) {
	if math.IsInf(d.F3, 0) || math.IsInf(d.F4, 0) {
		return d, fmt.Errorf("infinity result")
	}
	if math.IsNaN(d.F3) || math.IsNaN(d.F4) {
		return d, fmt.Errorf("NaN result")
	}
	return d, nil
}

func CalculateTest(a, b int16) Data {
	return Data{
		F1: int64(a + b),
		F2: int64(a * b),
		F3: math.Exp(float64(a)) * math.Exp(float64(b)),
		F4: math.Exp(float64(a)) * (-math.Exp(float64(b))),
	}
}
