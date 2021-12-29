package utils

import "strings"

// Escape a string by removing the line breaks.
// This can help to prevent log entry injections.
func EscapeString(input string) string {
	escapedString := strings.ReplaceAll(input, "\n", "")
	escapedString = strings.ReplaceAll(escapedString, "\r", "")
	return escapedString
}
