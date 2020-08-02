package main

import (
	"log"
	"math/big"
	"time"

	helper "github.com/cuken/ProjectEuler/Helper"
)

func main() {
	defer helper.TrackTime(time.Now(), "Problem 20")

	// Working with big int is not very much fun in golang.

	sum := big.NewInt(0)
	fac := big.NewInt(1)
	one := big.NewInt(1)
	zero := big.NewInt(0)
	ten := big.NewInt(10)
	s := big.NewInt(2)
	e := big.NewInt(100)

	for i := new(big.Int).Set(s); i.Cmp(e) < 0; i.Add(i, one) {
		fac.Mul(fac, i)
	}

	for i := new(big.Int).Set(fac); i.Cmp(zero) > 0; {
		v := fac.Mod(i, ten)
		sum.Add(sum, v)
		i.Div(i, ten)
		fac.Div(fac, ten)
	}

	log.Printf("The sum of all digits for 100! is: %v\n", sum)

}
