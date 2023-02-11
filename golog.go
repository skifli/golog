// Copyright skifli under the MIT license. ALl rights reserved.
// See '/LICENSE' for license information.
// SPDX-License-Identifier: MIT License.

// Package golog implements a logging infrastructure for Go.
// It focuses on performance, while providing a simple API.

package golog

import "os"

// Represents a log output.
// The zero value for Log is not ready for use, and must be populated.
type Log struct {
	writers   []*os.File    // The output to write the log messages to.
	formatter formatterType // The function which will format the log.
}

// Creates a new log output, and returns it.
// The formatter must follow the signature found in [formatterType].
func NewLog(writers []*os.File, formatter formatterType) *Log {
	return &Log{
		writers:   writers,
		formatter: formatter,
	}
}

// Adds a writer to the log.
func (log *Log) AddWriter(writer *os.File) {
	log.writers = append(log.writers, writer)
}

// Adds multiple writers to the log.
// Also see [Log.AddWriter].
func (log *Log) AddWriters(writers []*os.File) {
	log.writers = append(log.writers, writers...)
}

// Removes a writer from the log.
// It will remove as many instances of it as it can find.
func (log *Log) RemoveWriter(writer *os.File) {
	writersLen := len(log.writers)

	for i := 0; i < writersLen; i++ {
		if log.writers[i] == writer {
			log.writers[i] = log.writers[writersLen-1]
			log.writers = log.writers[:writersLen-1]

			i--
		}
	}
}

// Removes a writer from the log.
// It will remove as many instances of it as it can find.
// Also see [Log.RemoveWriter].
func (log *Log) RemoveWriters(writers []*os.File) {
	writersLen := len(log.writers)

	for i := 0; i < writersLen; i++ {
		for _, writer := range writers {
			if log.writers[i] == writer {
				log.writers[i] = log.writers[writersLen-1]
				log.writers = log.writers[:writersLen-1]

				i--
				break
			}
		}
	}
}

// Represents a logger.
// The zero value for Logger is ready for use.
type Logger struct {
	logs []*Log // Slice of log outputs.
}

// Creates a new logger, and returns it.
func NewLogger(logs []*Log) *Logger {
	return &Logger{logs: logs}
}

// Adds a log to the logger.
func (logger *Logger) AddLog(log *Log) {
	logger.logs = append(logger.logs, log)
}

// Adds multiple logs to the logger.
// Also see [Logger.AddLog]
func (logger *Logger) AddLogs(logs []*Log) {
	// Unpack the logs and pass themp to the append function
	// as variadic arguments.
	logger.logs = append(logger.logs, logs...)
}

// Removes a log from the logger.
// It will remove as many instances of it as it can find.
func (logger *Logger) RemoveLog(log *Log) {
	logsLen := len(logger.logs)

	for i := 0; i < logsLen; i++ {
		if logger.logs[i] == log {
			logger.logs[i] = logger.logs[logsLen-1]
			logger.logs = logger.logs[:logsLen-1]

			i--
		}
	}
}

// Removes multiple logs from the logger.
// It will remove as many instances of it as it can find.
// Also see [Logger.RemoveLog].
func (logger *Logger) RemoveLogs(logs []*Log) {
	logsLen := len(logger.logs)

	for i := 0; i < logsLen; i++ {
		for _, log := range logs {
			if logger.logs[i] == log {
				logger.logs[i] = logger.logs[logsLen-1]
				logger.logs = logger.logs[:logsLen-1]

				i--
				break
			}
		}
	}
}
