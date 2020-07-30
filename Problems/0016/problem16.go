package main

import (
	"log"
	"math/big"
	"strconv"
	"time"

	helper "github.com/cuken/ProjectEuler/Helper"
)

func main() {
	defer helper.TrackTime(time.Now(), "Problem 16")
	sum := 0

	n := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil).String()

	for _, c := range n {
		i, _ := strconv.Atoi(string(c))
		sum += i
	}

	log.Printf("The sum of 2^1000's digits is: %v", sum)

}
