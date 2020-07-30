package main

import (
	"log"
	"time"

	helper "github.com/cuken/ProjectEuler/Helper"
)

func main() {
	defer helper.TrackTime(time.Now(), "Problem 15")

	// We have to hard code the array lengths as make was being a jerk.
	gridSize := 20
	grid := [21][21]int{}

	for i := 0; i < gridSize; i++ {
		grid[i][gridSize] = 1
		grid[gridSize][i] = 1
	}

	for i := gridSize - 1; i >= 0; i-- {
		for j := gridSize - 1; j >= 0; j-- {
			grid[i][j] = grid[i+1][j] + grid[i][j+1]
		}
	}

	log.Printf("The total paths from the top left is: %v", grid[0][0])

}
