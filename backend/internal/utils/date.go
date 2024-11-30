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

func ParseDateNil(input *string) *time.Time {
	if input == nil {
		return nil
	}
	date, err := time.Parse("2006-01-02", *input)
	if err != nil {
		log.Println("error parsing date:", err)
		return nil
	}
	return &date
}

func StripTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func IsDateInRange(current, from, to time.Time) bool {
	return !current.Before(from) && !current.After(to)
}
