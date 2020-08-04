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

func PrimeFactorization(n int) (pfs map[int]int) {
	pfs = make(map[int]int)

	// Get the number of 2s that divide n
	for n%2 == 0 {
		if _, ok := pfs[2]; ok {
			pfs[2] += 1
		} else {
			pfs[2] = 1
		}
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			if _, ok := pfs[i]; ok {
				pfs[i] += 1
			} else {
				pfs[i] = 1
			}
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs[n] = 1
	}

	return
}

// PrimeFactors returns an int slice of all prime factors for n
func PrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

// SumOfPrimeFactors returns the sum of all primes up to n.
func SumOfPrimeFactors(num int, plist []int) (sum int) {
	n := num
	sum = 1
	p := plist[0]
	i := 0
	j := 0

	for p*p <= n && n > 1 && i < len(plist) {
		p = plist[i]
		i++
		if n%p == 0 {
			j = p * p
			n = n / p
			for n%p == 0 {
				j = j * p
				n = n / p
			}
			sum = sum * (j - 1) / (p - 1)
		}
	}

	if n > 1 {
		sum *= n + 1
	}

	return sum - num
}

// NumberOfDivisors calculates the number of divisors of a given number
func NumberOfDivisors(n int) int {
	pfs := PrimeFactorization(n)

	num := 1
	for _, exponents := range pfs {
		num *= (exponents + 1)
	}

	return num
}

// GetDivisors takes in a num and returns an int slice of all of its divisors and the count of divisors
func GetDivisors(num int) (divisors []int, count int) {
	t := math.Sqrt(float64(num))
	it := int(math.Round(t))
	is := []int{}

	for i := 1; i <= it; i++ {
		if it%i == 0 {
			if it/i == i {
				is = append(is, i)
			} else {
				is = append(is, i, it/i)
			}
		}
	}

	return is, len(is)
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
