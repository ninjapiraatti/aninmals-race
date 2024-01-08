package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/ninjapiraatti/aninmals-race/aninmals"
)

const numberOfContestants = 3
const raceLength = 20

func DisplayRace(aninmals []aninmals.Aninmal) {
	if numberOfContestants > 0 {
		fmt.Printf("\033[%dA", numberOfContestants)
	}

	for _, a := range aninmals {
		progressBar := a.Color + strings.Repeat("#", a.Progress) + strings.Repeat(".", raceLength-a.Progress) + "\033[0m"
		fmt.Printf("%-15s [%-10s]\n", a.Name, progressBar)
	}
}

func main() {
	fmt.Printf("\n\n\n\n")
	raceOver := false
	var mu sync.Mutex

	uniqueAninmals := make([]aninmals.Aninmal, numberOfContestants)
	for i := 0; i < numberOfContestants; {
		newAninmal := aninmals.Create()
		if !aninmals.IsDuplicate(newAninmal, uniqueAninmals) {
			uniqueAninmals[i] = newAninmal
			i++
		}
	}

	var wg sync.WaitGroup
	for i := range uniqueAninmals {
		wg.Add(1)
		go func(a *aninmals.Aninmal) {
			defer wg.Done()
			for {
				mu.Lock()
				if raceOver {
					mu.Unlock()
					break
				}
				a.Race()
				DisplayRace(uniqueAninmals)
				if a.Progress >= raceLength {
					raceOver = true
					fmt.Printf("%s won the race!\n", a.Name)
				}
				mu.Unlock()
				time.Sleep(100 * time.Millisecond)
			}
		}(&uniqueAninmals[i])
	}

	wg.Wait()
}
