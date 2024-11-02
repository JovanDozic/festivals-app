package utils

import (
	"log"
	"time"
)

func ParseDate(input string) time.Time {
	date, err := time.Parse("2006-01-02", input)
	if err != nil {
		log.Println("error parsing date:", err)
		return time.Time{}
	}
	return date
}
