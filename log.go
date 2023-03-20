// Copyright skifli under the MIT license. ALl rights reserved.
// See './LICENSE' for license information.
// SPDX-License-Identifier: MIT License.

// Package golog implements a logging infrastructure for Go.
// It focuses on performance, while providing a simple API.

package golog

import (
	"fmt"
	"os"
	"time"
)

// Logs the message with the fields to the log.
// Used internally.
func (logger *Logger) log(levelIndex levelIndexes, fields Fields, msg string) {
	now := time.Now()
	level := Levels[levelIndex]

	for _, log := range logger.logs {
		formattedMsg := log.formatter(now, level, msg, fields)

		for _, writer := range log.writers {
			_, err := writer.WriteString(formattedMsg + "\n")
			check(err)
		}
	}
}

// Logs a message, using FATAL as the log level.
// Then calls 'os.Exit(1)'.
func (logger *Logger) Fatal(msg string) {
	logger.log(FATAL, nil, msg)
	os.Exit(1)
}

// Logs a message with the specified fields, using FATAL as the log level.
// Then calls 'os.Exit(1)'.
func (logger *Logger) FatalFields(msg string, fields Fields) {
	logger.log(FATAL, fields, msg)
	os.Exit(1)
}

// Formats the message, using FATAL as the log level. Then calls 'os.Exit(1)'.
func (logger *Logger) Fatalf(msg string, format ...any) {
	logger.log(FATAL, nil, fmt.Sprintf(msg, format...))
	os.Exit(1)
}

// Formats the message, and then logs it with the specified fields, using FATAL
// as the log level. Then calls 'os.Exit(1)'.
func (logger *Logger) FatalFieldsf(msg string, fields Fields, format ...any) {
	logger.log(FATAL, fields, fmt.Sprintf(msg, format...))
	os.Exit(1)
}

// Logs the error, using ERROR as the log level. Then calls 'panic(err)'.
func (logger *Logger) Panic(err error) {
	logger.log(ERROR, nil, err.Error())
	panic(err)
}

// Logs the error with the specified fields, using ERROR as the log level.
// Then calls 'panic(err)'.
func (logger *Logger) PanicFields(err error, fields Fields) {
	logger.log(ERROR, fields, err.Error())
	panic(err)
}

// Formats the message, using ERROR as the log level. Then calls 'panic(err)'.
func (logger *Logger) Panicf(msg string, format ...any) {
	err := fmt.Sprintf(msg, format...)

	logger.log(ERROR, nil, err)
	panic(err)
}

// Formats the message, and then logs it with the specified fields, using ERROR
// as the log level. Then calls 'panic(err)'.
func (logger *Logger) PanicFieldsf(msg string, fields Fields, format ...any) {
	err := fmt.Sprintf(msg, format...)

	logger.log(ERROR, fields, err)
	panic(err)
}

// Logs a message, using ERROR as the log level.
func (logger *Logger) Error(msg string) {
	logger.log(ERROR, nil, msg)
}

// Logs a message with the specified fields, using ERROR as the log level.
func (logger *Logger) ErrorFields(msg string, fields Fields) {
	logger.log(ERROR, fields, msg)
}

// Formats the message, using ERROR as the log level.
func (logger *Logger) Errorf(msg string, format ...any) {
	logger.log(ERROR, nil, fmt.Sprintf(msg, format...))
}

// Formats the message, and then logs it with the specified fields, using ERROR
// as the log level.
func (logger *Logger) ErrorFieldsf(msg string, fields Fields, format ...any) {
	logger.log(ERROR, fields, fmt.Sprintf(msg, format...))
}

// Logs a message, using WARNING as the log level.
func (logger *Logger) Warning(msg string) {
	logger.log(WARNING, nil, msg)
}

// Logs a message with the specified fields, using WARNING as the log level.
func (logger *Logger) WarningFields(msg string, fields Fields) {
	logger.log(WARNING, fields, msg)
}

// Formats the message, using WARNING as the log level.
func (logger *Logger) Warningf(msg string, format ...any) {
	logger.log(WARNING, nil, fmt.Sprintf(msg, format...))
}

// Formats the message, and then logs it with the specified fields, using WARNING
// as the log level.
func (logger *Logger) WarningFieldsf(msg string, fields Fields, format ...any) {
	logger.log(WARNING, fields, fmt.Sprintf(msg, format...))
}

// Logs a message, using INFO as the log level.
func (logger *Logger) Info(msg string) {
	logger.log(INFO, nil, msg)
}

// Logs a message with the specified fields, using INFO as the log level.
func (logger *Logger) InfoFields(msg string, fields Fields) {
	logger.log(INFO, fields, msg)
}

// Formats the message, using INFO as the log level.
func (logger *Logger) Infof(msg string, format ...any) {
	logger.log(INFO, nil, fmt.Sprintf(msg, format...))
}

// Formats the message, and then logs it with the specified fields, using INFO
// as the log level.
func (logger *Logger) InfoFieldsf(msg string, fields Fields, format ...any) {
	logger.log(INFO, fields, fmt.Sprintf(msg, format...))
}

// Logs a message, using DEBUG as the log level.
func (logger *Logger) Debug(msg string) {
	logger.log(DEBUG, nil, msg)
}

// Logs a message with the specified fields, using DEBUG as the log level.
func (logger *Logger) DebugFields(msg string, fields Fields) {
	logger.log(DEBUG, fields, msg)
}

// Formats the message, using DEBUG as the log level.
func (logger *Logger) Debugf(msg string, format ...any) {
	logger.log(DEBUG, nil, fmt.Sprintf(msg, format...))
}

// Formats the message, and then logs it with the specified fields, using DEBUG
// as the log level.
func (logger *Logger) DebugFieldsf(msg string, fields Fields, format ...any) {
	logger.log(DEBUG, fields, fmt.Sprintf(msg, format...))
}
