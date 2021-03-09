package utils

import "strings"

// ErrorContains checks if the error message in out contains the text in
// want.
// This is safe when err is nil. Use an empty string for want if you want to
// test that err is nil.
func ErrorContains(err error, want string) bool {
	if err == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(err.Error(), want)
}
