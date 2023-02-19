// Copyright skifli under the MIT license. ALl rights reserved.
// See './LICENSE' for license information.
// SPDX-License-Identifier: MIT License.

package golog

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/goccy/go-json"
	"golang.org/x/term"
)

// Function signature of formatters.
// Used internally to reduce typing.
type formatterType = func(time.Time, string, string, Fields) string

// An example formatter for JSON.
// Returns a JSON-compliant representation of the log.
func FormatterJSON(
	now time.Time,
	level string,
	msg string,
	fields Fields,
) string {
	outputMap := Fields{"time": now.Unix(), "level": level, "msg": msg}

	for key, value := range fields {
		outputMap[key] = value
	}

	outputBytes, err := json.Marshal(outputMap)
	check(err)

	return string(outputBytes)
}

// An example formatter for humans.
// Returns a human-readable representation of the log.
func FormatterHuman(
	now time.Time,
	level string,
	msg string,
	fields Fields,
) string {
	// Get the width of the terminal
	width, _, err := term.GetSize(int(os.Stderr.Fd()))
	check(err)

	str := fmt.Sprintf(
		"%s[%s] %s%s%s %s",
		LevelColours[level],
		now.Format("02 Jan 2006 15:04:05 MST"),
		level,
		strings.Repeat(" ", LongestLevel-len(level)),
		LevelColours["RESET"],
		msg,
	)

	var fieldsBuilder strings.Builder

	for key, value := range fields {
		// Convert the fields into a string.
		fieldsBuilder.WriteString(key + "=" + fmt.Sprintf("%#v ", value))
	}

	fieldsStr := strings.TrimSuffix(fieldsBuilder.String(), " ")

	// Work out the amount of characters between 'str' and the end of the console.
	padding := width - len(str) - len(fieldsStr)

	if padding < 0 {
		padding = width - ((len(str) % width) - 9) - len(fieldsStr)

		// Console is *very* small
		if padding < 0 {
			padding = 1
		}

		fmt.Println(len(str), width, padding)
	}

	// Add the padding.
	str += strings.Repeat(" ", padding)

	// Add the fields.
	str += fieldsStr

	return str
}
