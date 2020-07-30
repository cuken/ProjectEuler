package main

import (
	"log"
	"time"

	helper "github.com/cuken/ProjectEuler/Helper"
)

func main() {

	// Track the time required to complete the solution
	defer helper.TrackTime(time.Now(), "Problem 2")

	// Start at 0
	sum := 0

	// Define our stopping point
	const max int = 4000000

	// Define our Fibonacci method
	f := helper.Fibonacci()
	v := f()

	for v < max {
		if v%2 == 0 {
			sum += v
		}
		v = f()
	}

	log.Printf("The sum of even Fibonacci digits is: %d\n", sum)

}
