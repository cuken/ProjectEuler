package main

import (
	"log"
	"time"

	helper "github.com/cuken/ProjectEuler/Helper"
)

func main() {

	defer helper.TrackTime(time.Now(), "Problem 7")
	p := helper.SieveofAtkin(1000000)

	log.Printf("The 10001 prime is %v", p[10000])

}
