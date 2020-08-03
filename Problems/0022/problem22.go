package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"time"

	helper "github.com/cuken/ProjectEuler/Helper"
)

func main() {
	defer helper.TrackTime(time.Now(), "Problem 22")

	sum := 0
	score := 0

	content, err := ioutil.ReadFile("p022_names.txt")
	if err != nil {
		log.Fatal("Unable to grab your names file yo!")
	}

	names := strings.Split(strings.Replace(string(content), "\"", "", -1), ",")
	sort.Strings(names)

	for i := 0; i < len(names); i++ {
		score = 0
		for n := 0; n < len(names[i]); n++ {
			//fmt.Printf("Letter: %s, Value: %v\n", string(names[i][n]), int(names[i][n]))
			score += int(names[i][n] - 64)
		}
		sum += score * (i + 1)
	}

	log.Printf("The sum of all name scores is: %v", sum)
}
