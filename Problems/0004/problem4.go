package main

import (
	"log"
	"time"

	helper "github.com/cuken/ProjectEuler/Helper"
)

func main() {

	x := 999
	y := 999
	r := 0

	defer helper.TrackTime(time.Now(), "Problem 4")

	// This will give us the first palindrome we see stepping backwards from one digit decrementing.

	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			if isNumPalindrome(i * j) {
				if i*j > r {
					r = i * j
					x = i
					y = j
				}
				continue
			}
		}
	}

	log.Printf("The largest palindrome for the product of two three digit numbers is: %v\nWhich is %v * %v", r, x, y)

}

// Used to determine if the number provided is a palindrome
func isNumPalindrome(num int) bool {
	var x int
	var y int = 0
	var t int = num

	// We'll extract the digits by dividing by it's position in the number
	// and using % 10 to find the remainder;
	for {
		x = num % 10
		y = y*10 + x
		num /= 10

		if num == 0 {
			break
		}
	}

	if t == y {
		return true
	}
	return false

}
