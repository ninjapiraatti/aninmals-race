package main

import (
	"log"

	"github.com/ninjapiraatti/aninmals-race/aninmals"
)

const numberOfContestants = 3

func main() {
	uniqueAninmals := make([]aninmals.Aninmal, numberOfContestants)
	for i := 0; i < numberOfContestants; {
		newAninmal := aninmals.Create()
		if !aninmals.IsDuplicate(newAninmal, uniqueAninmals) {
			uniqueAninmals[i] = newAninmal
			i++
		}
	}
	log.Println(uniqueAninmals)
}
