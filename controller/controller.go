package controller

import (
	"fmt"
	"log"
	"math"
	"worker/models"
)

type ControllerV1 struct {
	apiService        ApiService
	enableLog         bool
	ControllerVerSion string
}

type ApiService interface {
	Serve(cert, key string) error
}

func (c *ControllerV1) Serve(cert, key string) {
	err := c.apiService.Serve(cert, key)
	if err != nil {
		log.Fatal("Unable to load API service")
	}
}

func (c *ControllerV1) CalculateTest(a, b int16) (models.Data, error) {
	return models.Data{
		F1: int64(a + b),
		F2: int64(a * b),
		F3: math.Exp(float64(a)) * math.Exp(float64(b)),
		F4: math.Exp(float64(a)) * (-math.Exp(float64(b))),
	}, nil
}

func (c *ControllerV1) IsDebug() bool {
	return c.enableLog
}

func (c *ControllerV1) AddApiSerice(serv ApiService) {
	c.apiService = serv
}

func (c *ControllerV1) EnableLog() {
	c.enableLog = true
}

func NewController() *ControllerV1 {
	controller := ControllerV1{
		ControllerVerSion: "v1",
	}
	return &controller
}

func (c *ControllerV1) Log(msg string) {
	if c.enableLog {
		fmt.Println(msg)
	}
}
