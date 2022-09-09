package workers

import (
	"worker/models"
)

type IController interface {
	InjectWoker2(a, b int16)
}
type worker1 struct {
	controller IController
	a          chan int16
	b          chan int16
}

func (w *worker1) Start() {
	go func() {
		for {
			a := <-w.a
			b := <-w.b
			w.controller.InjectWoker2(a, b)
		}
	}()
}

func (w *worker1) HandleInput(a, b int16) {
	w.a <- a
	w.b <- b
}

func NewWorker1(controller IController) *worker1 {
	worker := &worker1{
		controller: controller,
		a:          make(chan int16, 1),
		b:          make(chan int16, 1),
	}
	return worker
}

type worker2 struct {
	controller IController
	a          chan int16
	b          chan int16
	output     chan models.Data
	e          chan error
}

func (w *worker2) Inject(a, b int16) {
	w.a <- a
	w.b <- b
}

func (w *worker2) Start() {
	go func() {
		for {
			datas, err := models.CalculateTest(<-w.a, <-w.b)
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
		a:          make(chan int16),
		b:          make(chan int16),
		output:     make(chan models.Data, 1),
		e:          make(chan error, 1),
	}
	return worker2
}
