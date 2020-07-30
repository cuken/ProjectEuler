package main

import (
	"log"
	"time"

	helper "github.com/cuken/ProjectEuler/Helper"
)

func main() {

	defer helper.TrackTime(time.Now(), "Problem 9")

	stop := true

	for a := 1; stop && a < 1000; a++ {
		for b := a + 1; stop && b < 1000; b++ {
			for c := 999; c > 3; c-- {
				if a+b+c == 1000 && isPythagorean(a, b, c) {
					s := a * b * c
					log.Printf("The largest product of 3 Pythagorean triples (%v,%v,%v) is: %v\n", a, b, c, s)
					stop = false
					break
				}
			}
		}
	}
}

func isPythagorean(a int, b int, c int) bool {
	as := a * a
	bs := b * b
	cs := c * c

	if as+bs == cs {
		return true
	}

	return false
}
