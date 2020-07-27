package helper

import (
	"log"
	"time"
)

// TrackTime prints out the time taken to execute the function.
func TrackTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s to complete", name, elapsed)
}

// Fibonacci all recurisve like
func Fibonacci() func() int {
	x, y := 1, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}
