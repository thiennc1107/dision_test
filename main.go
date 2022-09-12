package main

import (
	"fmt"
	"os"
	"worker/controller"
	"worker/handler"
	"worker/utils"
	"worker/workers"
)

func main() {
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
	ctrller.AddApiSerice(handler.NewAPiService(ctrller))
	ctrller.Serve("./cert/server.crt", "./cert/server.key")
}
