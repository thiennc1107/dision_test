package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"worker/controller"
	"worker/handler"
	"worker/utils"
	"worker/workers"
)

func main() {
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)
	ctrller := controller.NewController()
	ctrller.Worker1 = workers.NewWorker1(ctrller)
	ctrller.Worker2 = workers.NewWorker2(ctrller)
	var logger controller.Logger
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--version":
			fmt.Println(ctrller.ControllerVerSion)
			return
		case "--debug":
			fmt.Println("launching debug mode")
			logger = utils.NewLogger(true)
		}
	}
	if logger == nil {
		logger = utils.NewLogger(false)
	}
	ctrller.Logger = logger
	ctrller.CreateContext()
	ctrller.StartALl()
	go func() {
		<-sigChannel
		ctrller.StopAll()
		os.Exit(0)
	}()
	ctrller.AddApiSerice(handler.NewAPiService(ctrller))
	ctrller.Serve("./cert/server.crt", "./cert/server.key")
}
