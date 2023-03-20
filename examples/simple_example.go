package main

import (
	"errors"
	"os"

	"github.com/skifli/golog"
)

// Create the logger with the console log already configured.
var logger = golog.NewLogger([]*golog.Log{
	golog.NewLog(
		[]*os.File{
			os.Stderr,
		},
		golog.FormatterHuman,
	),
})

func main() {
	logFile, err := os.OpenFile("simple_example.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)

	if err != nil {
		logger.Panic(err)
	}

	// Add the file log to the logger.
	logger.AddLog(golog.NewLog(
		[]*os.File{
			logFile,
		},
		golog.FormatterJSON,
	))

	// Will be run after the 'logger.Panic' call.
	defer func() {
		if r := recover(); r != nil {
			logger.Infof("recovered from the panic '%v'", r)
			logger.Fatal("and this is a fatal error")
		}
	}()

	logger.Debugf("this is a %s message", "debug")
	logger.Info("some useful text")
	logger.WarningFieldsf(
		"this is a %s message - you can even add fields",
		golog.Fields{"service": "golog"},
		"warning",
	)
	logger.Error("this is an error")
	logger.Panic(errors.New("this is a panic on an error"))
}
