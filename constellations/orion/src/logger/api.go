package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

var PROD_LOG_FILE_NAME string = "service.log"

type Fields = log.Fields

var standardFields = Fields{
	"app": "orion",
}

func SetupDisabled() {
	log.SetOutput(ioutil.Discard)
}

func SetupTest() {
	log.New()
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
}

func SetupDev() {
	log.New()
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stderr)
}

func SetupProd() error {
	log.New()
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	// Log to local file
	file, err := os.OpenFile("service.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		return err
	}
	log.SetOutput(file)
	log.Println("Service log re-opened")
	return nil
}

func ClearLogFile(filePath string) error {
	err := os.Remove(PROD_LOG_FILE_NAME)
	if err != nil {
		log.Printf("Error deleting file %s (%w)", filePath, err)
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
		Error(fmt.Sprintf("%s (%v)", message, err))
	// Error(fmt.Spritnf("%s (%+v)", message, err)) // <- use if want to see stacktrace
	// Error(fmt.Sprintf("%s (%w)", message, err)) // <- displays object memory address
}

func CombineFields(fieldsA, fieldsB Fields) Fields {
	for k, v := range fieldsB {
		fieldsA[k] = v
	}
	return fieldsA
}
