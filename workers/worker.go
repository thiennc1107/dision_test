package workers

import (
	"context"
	"worker/models"
)

type IController interface {
	InjectWorker2(input models.Input)
	GetContext() *context.Context
}
type worker1 struct {
	controller IController
	input      chan models.Input
	ctx        *context.Context
}

func (w *worker1) Start() {
	go func() {
		for {
			input := <-w.input
			w.controller.InjectWorker2(input)
		}
	}()
}

func (w *worker1) HandleInput(a, b int16) {
	input := models.Input{
		A: a,
		B: b,
	}
	w.input <- input
}

func NewWorker1(controller IController) *worker1 {
	worker := &worker1{
		controller: controller,
		input:      make(chan models.Input),
		ctx:        controller.GetContext(),
	}
	return worker
}

type worker2 struct {
	controller IController
	input      chan models.Input
	output     chan models.Data
	e          chan error
	ctx        *context.Context
}

func (w *worker2) Inject(input models.Input) {
	w.input <- input
}

func (w *worker2) Start() {
	go func() {
		for {
			input := <-w.input
			datas, err := models.CalculateTest(input.A, input.B)
			if err != nil {
				w.e <- err
				continue
			}
			err = datas.CheckInvalidResult()
			if err != nil {
				w.e <- err
				continue
			}
			w.output <- datas
		}
	}()
}

func (w *worker2) GetOutPut() (models.Data, error) {
	select {
	case err := <-w.e:
		return models.Data{}, err
	case data := <-w.output:
		return data, nil
	}
}

func NewWorker2(controller IController) *worker2 {
	worker2 := &worker2{
		controller: controller,
		// a:          make(chan int16),
		// b:          make(chan int16),
		input:  make(chan models.Input),
		output: make(chan models.Data),
		e:      make(chan error),
		ctx:    controller.GetContext(),
	}
	return worker2
}
