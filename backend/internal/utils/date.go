package utils

import (
	"log"
	"time"
)

// Parses date string (yyyy-mm-dd) to time.Time
func ParseDate(input string) time.Time {
	date, err := time.Parse("2006-01-02", input)
	if err != nil {
		log.Println("error parsing date:", err)
		return time.Time{}
	}
	return date
}

// Returns true if the input string is a valid date string (yyyy-mm-dd)
func ValidateDate(input string) bool {
	_, err := time.Parse("2006-01-02", input)
	return err == nil
}
