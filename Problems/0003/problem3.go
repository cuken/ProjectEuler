package main

import (
	"log"
	"time"

	helper "github.com/cuken/ProjectEuler/Helper"
)

func main() {

	defer helper.TrackTime(time.Now(), "Problem 3")
	p := 600851475143
	n := p
	p1 := []int{}

	for n%2 == 0 {
		p1 = append(p1, 2)
		n = n / 2
	}

	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			p1 = append(p1, i)
			n = n / i
		}
	}

	if n > 2 {
		p1 = append(p1, n)
	}

	// Can't use Sieve, :(
	// r := helper.SieveofEratosthenes(p)

	log.Printf("The largest prime number for %v is %v\n Factorizations are: %v", p, n, p1)
}
