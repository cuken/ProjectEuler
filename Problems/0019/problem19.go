package main

import (
	"log"
	"time"

	helper "github.com/cuken/ProjectEuler/Helper"
)

func main() {
	defer helper.TrackTime(time.Now(), "Problem 19")

	// 1 Jan 1900 -> monday
	year := 1901
	month := 1
	day := 1
	dayMax := 0
	dow := 2
	isLeap := false
	sunCount := 0

	for year < 2001 {
		if year%4 == 0 {
			if year%100 == 0 && year%400 == 0 {
				isLeap = true
			} else if year%100 == 0 {
				isLeap = false
			} else {
				isLeap = true
			}
		} else {
			isLeap = false
		}
		for month <= 12 {
			switch month {
			case 1, 3, 5, 7, 8, 10, 12:
				dayMax = 31
			case 2:
				if isLeap {
					dayMax = 29
				} else {
					dayMax = 28
				}
			case 4, 6, 9, 11:
				dayMax = 30
			}
			for day <= dayMax {
				if dow == 7 {
					if day == 1 {
						sunCount++
					}
					dow = 1
				} else {
					dow++
				}
				day++
			}
			month++
			day = 1
		}
		year++
		month = 1
	}

	log.Printf("The total number of sundays between 1901 to 2000: %v", sunCount)

}
