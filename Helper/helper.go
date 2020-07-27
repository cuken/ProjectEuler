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
