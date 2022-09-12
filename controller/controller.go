package controller

import (
	"context"
	"fmt"
	"log"
	"time"
	"worker/models"
)

type ControllerV1 struct {
	apiService        ApiService
	enableLog         bool
	ControllerVerSion string
	Worker1           IWorker1
	Worker2           IWorker2
	ctx               context.Context
	cancel            context.CancelFunc
	timeOut           int64
}

type ApiService interface {
	Serve(cert, key string) error
}

type IWorker1 interface {
	HandleInput(a, b int16)
	Start()
	Stop()
}

type IWorker2 interface {
	Inject(input models.Input)
	GetOutPut() (models.Data, error)
	Start()
	Stop()
}

func (c *ControllerV1) GetContext() context.Context {
	return c.ctx
}

func (c *ControllerV1) createContext() {
	context, cancel := context.WithTimeout(context.Background(),
		time.Duration(c.timeOut)*time.Second)
	c.ctx = context
	c.cancel = cancel
}

func (c *ControllerV1) Serve(cert, key string) {
	err := c.apiService.Serve(cert, key)
	if err != nil {
		log.Fatal("Unable to load API service")
	}
}

func (c *ControllerV1) CalculateTest(a, b int16) (models.Data, error) {
	c.createContext()
	c.StartALl()
	c.Worker1.HandleInput(a, b)
	data, err := c.Worker2.GetOutPut()
	// TODO: create time out response
	c.StopAll()
	if err != nil {
		return models.Data{}, err
	}
	return data, nil
}

func (c *ControllerV1) StopAll() {
	c.Worker1.Stop()
	c.Worker2.Stop()
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

func (c *ControllerV1) InjectWorker2(input models.Input) {
	c.Worker2.Inject(input)
}

func (c *ControllerV1) StartALl() {
	c.Worker1.Start()
	c.Worker2.Start()
}

func NewController() *ControllerV1 {
	controller := ControllerV1{
		ControllerVerSion: "v1",
		timeOut:           5,
	}
	return &controller
}
