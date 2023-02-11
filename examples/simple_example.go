package main

import (
	"errors"
	"os"

	"github.com/skifli/golog"
)

func main() {
	logFile, err := os.OpenFile("simple_example.log", os.O_CREATE|os.O_APPEND, 0777)

	if err != nil {
		panic(err)
	}

	// Create the logger with the logs already configured.
	// Logs can also be added / removed later.
	logger := golog.NewLogger([]*golog.Log{
		golog.NewLog(
			[]*os.File{
				os.Stderr,
			},
			golog.FormatterHuman,
		),
		golog.NewLog(
			[]*os.File{
				logFile,
			},
			golog.FormatterJSON,
		),
	})

	defer func() {
		logger.Info("Recovered from the panic.", nil)
		logger.Fatal("And this is a fatal error. Goodbye.", nil)
	}()

	logger.Debugf("This is a %s message.", nil, "debug")
	logger.Info("Some useful text...", nil)
	logger.Warningf(
		"This is a %s message - you can even add fields.",
		golog.Fields{"service": "golog"},
		"warning",
	)
	logger.Error("This is an error.", nil)
	logger.Panic(errors.New("This is a panic on an error"), nil)
}
