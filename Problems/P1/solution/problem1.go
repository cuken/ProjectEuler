// This is the solution to Problem 1 for Project Euler.

package main

import (
	"log"
	"time"

	helper "github.com/cuken/ProjectEuler/Helper"
)

// Define our entry point.
func main() {

	// Track the time required to complete the solution
	defer helper.TrackTime(time.Now(), "Problem 1")

	// Start at 0
	sum := 0

	// Start at 1 and go through
	for i := 1; i < 1000; i++ {

		// Check if I is divisible by 3 or 5;
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	log.Printf("The sum of all digits: %d\n", sum)
}
