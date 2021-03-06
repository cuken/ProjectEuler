package main

import (
	"log"
	"math"
	"time"

	helper "github.com/cuken/ProjectEuler/Helper"
)

func main() {

	defer helper.TrackTime(time.Now(), "Problem 21")
	sum := 0
	// plist := helper.SieveofAtkin(int(math.Sqrt(10001)))
	plist := helper.SieveofEratosthenes(int(math.Sqrt(10001)))

	for i := 2; i <= 10000; i++ {
		factors := helper.SumOfPrimeFactors(i, plist)
		if factors > i && factors <= 10000 {
			factorsj := helper.SumOfPrimeFactors(factors, plist)
			if factorsj == i {
				sum += i + factors
			}
		}
	}

	log.Printf("The sum of all amicable pairs below 10000 is: %v", sum)
}

func main1() {
	defer helper.TrackTime(time.Now(), "Problem 21")

	sum := 0

	for i := 2; i <= 10000; i++ {
		faci := sumOfFactors(i)
		if faci > i && faci <= 10000 {
			facj := sumOfFactors(faci)
			if facj == i {
				sum += i + faci
			}
		}
	}

	log.Printf("The sum of all amicable pairs below 10000 is: %v", sum)

}

func sumOfFactors(num int) (sum int) {
	sqrtF := math.Sqrt(float64(num))
	sqrt := int(sqrtF)
	sum = 1

	if num == sqrt*sqrt {
		sum += sqrt
		sqrt--
	}

	for i := 2; i < sqrt; i++ {
		if num%i == 0 {
			sum = sum + i + (num / i)
		}
	}
	return sum
}
