package main

import (
	"log"
	"math"
	"time"

	helper "github.com/cuken/ProjectEuler/Helper"
)

func main() {
	defer helper.TrackTime(time.Now(), "Problem 23")

	sum := 0
	max := 28123 // From the problem.
	plist := helper.SieveofEratosthenes(int(math.Sqrt(float64(max))))
	var alist []int

	for i := 2; i < max; i++ { // find all abundant numbers
		if helper.SumOfPrimeFactors(i, plist) > i {
			alist = append(alist, i)
		}
	}

	canBeAbund := make([]bool, max+1)

	// loop through all abudant numbers
	for i := 0; i < len(alist); i++ {
		for j := i; j < len(alist); j++ {
			if alist[i]+alist[j] <= max {
				canBeAbund[alist[i]+alist[j]] = true
			} else {
				break
			}
		}
	}

	//Look for all of the false abundant numbers and sum them.
	for i := 0; i <= max; i++ {
		if !canBeAbund[i] {
			sum += i
		}
	}

	log.Printf("The sum of all numbers that cannot be the sum of two abundant numbers is: %v", sum)

}
