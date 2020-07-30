package main

import (
	"log"
	"time"

	helper "github.com/cuken/ProjectEuler/Helper"
)

func main() {
	defer helper.TrackTime(time.Now(), "Problem 12")
	num := 1
	t := 1

	for {
		t = findTriangleSum(num)
		x := helper.NumberOfDivisors(t)

		if x > 500 {
			break
		}
		num++
	}

	log.Printf("The value of the first triangle number to have over five hundred divisors is: %v", t)

}

func findTriangleSum(num int) int {
	return num * (num + 1) / 2 // way more efficient: https://www.geeksforgeeks.org/program-find-sum-first-n-natural-numbers/
}
