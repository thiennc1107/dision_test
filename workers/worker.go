package workers

import (
	"context"
	"fmt"
	"sync"
	"worker/models"
	"worker/utils"
)

type IController interface {
	InjectWorker2(input models.Input)
	GetContext() context.Context
	Log(msg string)
}
type worker1 struct {
	controller IController
	input      chan models.Input
	close      chan bool
}

var wg sync.WaitGroup

func (w *worker1) Start() {
	w.controller.Log("Starting worker 1")
	go func() {
		var input models.Input
	outer:
		for {
			w.controller.Log("worker 1 pending")
			select {
			case input = <-w.input:
				w.controller.Log("worker 1 working")
			case <-w.close:
				w.controller.Log("worker 1 stopping")
				break outer
			}
			w.controller.InjectWorker2(input)
		}
		w.controller.Log("worker 1 stopped")
		wg.Done()
	}()
}

func (w *worker1) Stop() {
	w.controller.Log("sending signal to close worker 1")
	wg.Add(1)
	w.close <- true
	wg.Wait()
}

func (w *worker1) HandleInput(a, b int16) {
	// time.Sleep(3 * time.Second)
	input := models.Input{
		A: a,
		B: b,
	}
	context := w.controller.GetContext()
	select {
	case <-context.Done():
		w.controller.Log("Handle Input cancelled")
	case w.input <- input:
		w.controller.Log("Input handled")
	}
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
	context := w.controller.GetContext()
	select {
	case <-context.Done():
		w.controller.Log("Inject cancelled")
	case w.input <- input:
		w.controller.Log("Data transferred")
	}
}

func (w *worker2) Start() {
	w.controller.Log("Starting worker 2")
	go func() {
	outer:
		for {
			var input models.Input
			w.controller.Log("worker 2 pending")
			select {
			case input = <-w.input:
				w.controller.Log("worker 2 working")
			case <-w.close:
				w.controller.Log("stopping worker 2")
				break outer
			}
			context := w.controller.GetContext()
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
		w.controller.Log("worker 2 stopped")
		wg.Done()
	}()
}

func (w *worker2) Stop() {
	w.controller.Log("sending signal to close worker 2")
	wg.Add(1)
	w.close <- true
	wg.Wait()
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
