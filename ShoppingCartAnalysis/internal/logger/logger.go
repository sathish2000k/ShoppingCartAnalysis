package logger

import (
	"io"
	"log"
	"os"
)

var (
	LoggerInfo *log.Logger
	LoggerError *log.Logger
)

func InitLogger() {
	serviceLogFile, err := os.OpenFile("shopping-cart-analysis.log", os.O_CREATE| os.O_WRONLY | os.O_APPEND ,0666)
	if err != nil {
		log.Fatalln("Failed to open log file: ", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, serviceLogFile)


	LoggerInfo = log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	LoggerError = log.New(multiWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}