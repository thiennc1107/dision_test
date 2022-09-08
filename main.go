package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"worker/controller"
	"worker/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	controller.RegisterController()
	if len(os.Args) > 1 {
		if os.Args[1] == "--version" {
			controller.ListController()
			return
		}
	}
	r := gin.New()

	server := http.Server{
		Addr:    ":1234",
		Handler: r,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS13,
		},
	}

	r.GET("/api/:version/test", handler.TestHandler)

	err := server.ListenAndServeTLS("./cert/server.crt", "./cert/server.key")
	if err != nil {
		log.Println(err)
		log.Panicln("Failed init server")
	}
}
