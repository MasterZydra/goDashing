package utils

import (
	"html"
	"strings"
)

// Escape a string by removing the line breaks and escaping HTML.
// This can help to prevent log entry injections.
func EscapeString(input string) string {
	escapedString := strings.ReplaceAll(input, "\n", "")
	escapedString = strings.ReplaceAll(escapedString, "\r", "")
	escapedString = html.EscapeString(escapedString)
	return escapedString
}
