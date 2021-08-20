package utils

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func InitLogger(LogFilePath string) {
	stdout := os.Stdout
	if len(LogFilePath) > 0 {
		var err error
		stdout, err = os.OpenFile("LogFilePath", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
	}

	InfoLogger = log.New(stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Warn(msg string) {
	WarningLogger.Println(msg)
}
