package utils

import (
	"log"
	"os"
)

const path = "logs.txt"

type logger struct {
	logger *log.Logger
}

func NewLogger(debug bool) *logger {
	loggerCore := log.Default()
	if !debug {
		file, err := os.OpenFile(path, os.O_APPEND|
			os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		loggerCore.SetOutput(file)
	}
	l := &logger{}
	l.logger = loggerCore
	return l
}

func (l *logger) Log(msg string) {
	l.logger.Println(msg)
}
