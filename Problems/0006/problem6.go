package main

import (
	"log"
	"time"

	helper "github.com/cuken/ProjectEuler/Helper"
)

func main() {

	defer helper.TrackTime(time.Now(), "Problem 6")

	fs := 0
	ss := 0

	for i := 1; i <= 100; i++ {
		fs += i * i
		ss += i
	}

	ss = ss * ss

	r := ss - fs

	log.Printf("The difference between the sum of squares is %v", r)

}
