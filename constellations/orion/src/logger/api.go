package logger

import (
	"fmt"
	// "os"
	log "github.com/sirupsen/logrus"
)

var LOG_FILE_NAME string = "service.log"

type Fields = log.Fields

var standardFields = Fields {
	"appName":  "orion",
}

func SetupTest() {
	logrus.New()
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)
}

func SetupDev() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)
}

func SetupProd() error {
	log.SetFormatter(&log.JSONFormatter{})

	// Log to local file
	file, err := os.OpenFile(LOG_FILE_NAME, os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
    if err != nil {
		fmt.Printf("error opening file: %v", err)
		return err
    }
	log.SetOutput(file)
	log.Println("Service log re-opened")

	log.SetLevel(log.InfoLevel)
	return nil
}

func ClearLogFile() error {
	err := os.Remove(LOG_FILE_NAME) 
    if err != nil { 
		log.Printf("Error deleting file %s (%w)", LOG_FILE_NAME, err)
		return err
	} 
	return nil
}

func Debug(message string, fields Fields) {
	log.WithFields(standardFields).
		WithFields(fields).
		Debug(message)
}

func Info(message string, fields Fields) {
	log.WithFields(standardFields).
		WithFields(fields).
		Info(message)
}

func Message(message string) {
	Info(message, Fields{})
}

func Error(message string, err error, fields Fields) {
	log.WithFields(standardFields).
		WithFields(fields).
		Error(fmt.Sprintf("%s (%w)", message, err))
}