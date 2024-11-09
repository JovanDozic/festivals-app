package utils

import (
	"regexp"
	"time"
)

// Returns true if the input string is a valid email address
func IsEmailValid(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(pattern, email)
	if err != nil {
		return false
	}
	return matched
}

// Returns true if the input string is a valid date string (yyyy-mm-dd)
func IsDateValid(input string) bool {
	_, err := time.Parse("2006-01-02", input)
	return err == nil
}
