package controller

import (
	"fmt"
	"log"
	"worker/models"
)

type ControllerV1 struct {
	apiService        ApiService
	enableLog         bool
	ControllerVerSion string
	Worker1           IWorker1
	Worker2           IWorker2
}

type ApiService interface {
	Serve(cert, key string) error
}

type IWorker1 interface {
	HandleInput(a, b int16)
	Start()
}

type IWorker2 interface {
	Inject(a, b int16)
	GetOutPut() (models.Data, error)
	Start()
}

func (c *ControllerV1) Serve(cert, key string) {
	err := c.apiService.Serve(cert, key)
	if err != nil {
		log.Fatal("Unable to load API service")
	}
}

func (c *ControllerV1) CalculateTest(a, b int16) (models.Data, error) {
	c.Worker1.HandleInput(a, b)
	data, err := c.Worker2.GetOutPut()
	if err != nil {
		return models.Data{}, err
	}
	return data, nil
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

func (c *ControllerV1) Log(msg string) {
	if c.enableLog {
		fmt.Println(msg)
	}
}

func (c *ControllerV1) InjectWorker2(a, b int16) {
	c.Worker2.Inject(a, b)
}

func (c *ControllerV1) StartALl() {
	c.Worker1.Start()
	c.Worker2.Start()
}

func NewController() *ControllerV1 {
	controller := ControllerV1{
		ControllerVerSion: "v1",
	}
	return &controller
}
