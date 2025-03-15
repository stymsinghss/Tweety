package utils

import "regexp"

// Common error messages used throughout the app
var (
	// RxUsername -> regular expression for username validation
	RxUsername = regexp.MustCompile("^[a-zA-Z][[a-zA-Z0-9_-]{0,17}$")

	// RxEmail -> regular expression for email validation
	RxEmail = regexp.MustCompile("^[^\\s@]+@[^\\s@]+\\.[^\\s@]+$")
)
