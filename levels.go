// Copyright skifli under the MIT license. ALl rights reserved.
// See './LICENSE' for license information.
// SPDX-License-Identifier: MIT License.

package golog

// Represents a level's index.
// Used internally.
type levelIndexes int

// Enum to represent a level's index.
const (
	FATAL levelIndexes = iota
	ERROR
	WARNING
	INFO
	DEBUG
)

// Level names.
var Levels = []string{
	"FATAL",
	"ERROR",
	"WARNING",
	"INFO",
	"DEBUG",
}

var LevelColours = map[string]string{
	"FATAL":   "\u001b[35m", // Magenta
	"ERROR":   "\u001b[31m", // Red
	"WARNING": "\u001b[33m", // Yellow
	"INFO":    "\u001b[34m", // Blue
	"DEBUG":   "",

	"RESET": "\u001b[0m",
}

var LongestLevel = func() int {
	longestLevelLength := 0

	for _, levelName := range Levels {
		levelLength := len(levelName)

		if levelLength > longestLevelLength {
			longestLevelLength = levelLength
		}
	}

	return longestLevelLength
}()
