// Copyright skifli under the MIT license. ALl rights reserved.
// See '/LICENSE' for license information.
// SPDX-License-Identifier: MIT License.

package golog

// Represents the fields arguments used when logging.
type Fields = map[string]any

// Checks if an error was returned, if so panic.
// Used internally.
func check(err error) {
	if err != nil {
		panic(err)
	}
}
