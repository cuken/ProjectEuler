package main

import (
	"log"
	"time"

	helper "github.com/cuken/ProjectEuler/Helper"
)

func main() {
	defer helper.TrackTime(time.Now(), "Problem 14")
	maxCount := 0
	maxNum := 0

	for i := 999999; i > 777777; i-- {
		count := 0
		s := i
		so := s
		for s > 1 {
			s = numDance(s)
			count++
		}

		if count > maxCount {
			maxNum = so
			maxCount = count
		}
	}

	log.Printf("The longest chain was %v, for the number %v", maxCount, maxNum)
}

func numDance(n int) (res int) {
	// n → n/2 (n is even)
	// n → 3n + 1 (n is odd)
	if n%2 == 0 {
		return n / 2
	}

	return (n * 3) + 1
}
