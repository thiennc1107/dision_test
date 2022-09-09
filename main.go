package main

import (
	"fmt"
	"os"
	"worker/controller"
	"worker/handler"
)

func main() {
	ctrller := controller.NewController()
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--version":
			fmt.Println(ctrller.ControllerVerSion)
			return
		case "--debug":
			ctrller.EnableLog()
		}
	}
	ctrller.AddApiSerice(handler.NewAPiService(ctrller))
	ctrller.Serve("./cert/server.crt", "./cert/server.key")
}
