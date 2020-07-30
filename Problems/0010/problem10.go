package main

import (
	"log"
	"time"

	helper "github.com/cuken/ProjectEuler/Helper"
)

func main() {
	defer helper.TrackTime(time.Now(), "Problem 10")

	p := helper.SieveofAtkin(2000000)
	s := 0

	for _, i := range p {
		s += i
	}

	log.Printf("The sum of all primes up to 2 million is: %v", s)
}
