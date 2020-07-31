package main

import (
	"log"
	"time"

	helper "github.com/cuken/ProjectEuler/Helper"
)

//			    one two three four five six seven eight nine
//				ten eleven twelve thirteen fourteen fifteen sixteen seventeen eighteen nineteen
//				twenty thirty fourty fifty sixty seventy eighty ninety
var lC1 = []int{0, 3, 3, 5, 4, 4, 3, 5, 5, 4}
var lC10e = []int{3, 6, 6, 8, 8, 7, 7, 9, 8, 8}
var lC10 = []int{0, 0, 6, 6, 5, 5, 5, 7, 6, 6}

func main() {
	defer helper.TrackTime(time.Now(), "Problem 17")

	s := 0

	// Start at 1 and go to 1000
	for i := 1; i < 1001; i++ {
		s += returnLetterCount(i)
	}

	log.Printf("The sum of all letters in the words of numbers from 1 to 1000 is %v", s)
}

func returnLetterCount(num int) int {
	// Edge cases;
	// 11, 12, and then all the 1* have special numbers.
	// everything else follows a 10's term and then the single digit.
	c := 0

	if num == 1000 {
		return 11
	}

	// check if number contains hundreds?
	if num > 99 {
		if num%100 > 0 {
			c += 3 // <-- add 3 for and
		}
		h := num / 100
		c += lC1[h] + 7 // number + hundred in text
		num -= (100 * h)
	}
	if num > 9 && num < 20 { // Special 10->19 case
		num -= 10
		c += lC10e[num]
		num = 0
	} else if num > 19 {
		h := num / 10
		c += lC10[h]
		num -= (10 * h)
	}

	c += lC1[num]

	return c
}
