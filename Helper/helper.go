package helper

import (
	"log"
	"math"
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

// SieveofEratosthenes returns a []int slice that has a list of all primes up to (not including) n.
func SieveofEratosthenes(n int) []int {
	f := make([]bool, n)
	var r []int

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if f[i] == false {
			for j := i * i; j < n; j += i {
				f[j] = true
			}
		}
	}
	for i := 2; i < n; i++ {
		if f[i] == false {
			r = append(r, i)
		}
	}

	return r
}
