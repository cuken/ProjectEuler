package main

import (
	"log"
	"time"

	helper "github.com/cuken/ProjectEuler/Helper"
)

func main() {

	defer helper.TrackTime(time.Now(), "Problem 5")

	// Start at the maximum number we'll divide by
	x := 20
	s := false
	b := false
	// We'll run it through the gauntlet!

	for !s {
		for i := 2; i < 20; i++ {
			if x%i != 0 {
				b = false
				// Skip by 10 as we need the last digit to be 0 to work with / 10 evenly;
				x += 10
				break
			} else {
				b = true
			}
		}
		if b {
			s = true
		}
	}

	log.Printf("The smallest number that cleanly divides into 1...20 is %v", x)
}
