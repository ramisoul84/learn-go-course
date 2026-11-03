// Package greeting builds friendly greetings.
package greeting

import "strings"

// Hello returns a greeting. It starts with a capital letter, so it's exported.
func Hello(name string) string {
	return "Hello, " + capitalize(name) + "!"
}

// capitalize starts with a small letter, so only this package can use it.
func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
