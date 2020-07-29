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

// SieveofAtkin will find all primes up to a number N and return the primes in a slice.
func SieveofAtkin(N int) []int {
	var x, y, n int
	nsqrt := math.Sqrt(float64(N))

	isPrime := make([]bool, N)

	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4*(x*x) + y*y
			if n <= N && (n%12 == 1 || n%12 == 5) {
				isPrime[n] = !isPrime[n]
			}
			n = 3*(x*x) + y*y
			if n <= N && n%12 == 7 {
				isPrime[n] = !isPrime[n]
			}
			n = 3*(x*x) - y*y
			if x > y && n <= N && n%12 == 11 {
				isPrime[n] = !isPrime[n]
			}
		}
	}

	for n = 5; float64(n) <= nsqrt; n++ {
		if isPrime[n] {
			for y = n * n; y < N; y += n * n {
				isPrime[y] = false
			}
		}
	}

	isPrime[2] = true
	isPrime[3] = true

	primes := make([]int, 0, 1270606)
	for x = 0; x < len(isPrime)-1; x++ {
		if isPrime[x] {
			primes = append(primes, x)
		}
	}

	return primes
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
