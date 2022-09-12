package workers

import (
	"context"
	"fmt"
	"worker/models"
	"worker/utils"
)

type IController interface {
	InjectWorker2(input models.Input)
	GetContext() context.Context
}
type worker1 struct {
	controller IController
	input      chan models.Input
	close      chan bool
}

func (w *worker1) Start() {
	go func() {
		var input models.Input
	outer:
		for {
			fmt.Println("worker 1 pending")
			select {
			case input = <-w.input:
				fmt.Println("worker 1 working")
			case <-w.close:
				fmt.Println("worker 1 stopping")
				break outer
			}
			w.controller.InjectWorker2(input)
		}
		fmt.Println("worker 1 stopped")
	}()
}

func (w *worker1) Stop() {
	fmt.Println("sending signal to close worker 1")
	w.close <- true
}

func (w *worker1) HandleInput(a, b int16) {
	// time.Sleep(5 * time.Second)s
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
		close:      make(chan bool),
	}
	return worker
}

type worker2 struct {
	controller IController
	input      chan models.Input
	output     chan models.Data
	e          chan error
	close      chan bool
}

func (w *worker2) Inject(input models.Input) {
	w.input <- input
}

func (w *worker2) Start() {
	go func() {
	outer:
		for {
			context := w.controller.GetContext()
			var input models.Input
			fmt.Println("worker 2 pending")
			select {
			case input = <-w.input:
				fmt.Println("worker 2 working")
			case <-w.close:
				fmt.Println("stopping worker 2")
				break outer
			}
			datas, err := models.CalculateTest(input.A, input.B).
				CheckInvalidResult()
			if err != nil {
				select {
				case w.e <- err:
					continue
				case <-context.Done():
					continue
				}
			}
			select {
			case w.output <- datas:
				continue
			case <-context.Done():
				continue
			}
		}
		fmt.Println("worker 2 stopped")
	}()
}

func (w *worker2) Stop() {
	fmt.Println("sending signal to close worker 2")
	w.close <- true
}

func (w *worker2) GetOutPut() (models.Data, error) {
	context := w.controller.GetContext()
	select {
	case err := <-w.e:
		return models.Data{}, err
	case data := <-w.output:
		return data, nil
	case <-context.Done():
		return models.Data{}, fmt.Errorf(utils.RequestTimedOut)
	}
}

func NewWorker2(controller IController) *worker2 {
	worker2 := &worker2{
		controller: controller,
		close:      make(chan bool),
		input:      make(chan models.Input),
		output:     make(chan models.Data),
		e:          make(chan error),
	}
	return worker2
}
