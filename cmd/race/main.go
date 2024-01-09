package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/ninjapiraatti/aninmals-race/aninmals"
)

const numberOfContestants = 3
const raceLength = 40

func DisplayRace(ctx context.Context, aninmals []aninmals.Aninmal) {
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if numberOfContestants > 0 {
				fmt.Printf("\033[%dA", numberOfContestants)
			}

			for _, a := range aninmals {
				progressBar := strings.Repeat("#", a.Progress) + strings.Repeat(".", raceLength-a.Progress)
				fmt.Printf("%s%-30s %-2s [%-35s]\033[0m\n", a.Color, a.Name, a.ProgressFromRedis, progressBar)
			}
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	println("\n\n\n")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6382",
		Password: "",
		DB:       0,
	})

	uniqueAninmals := make([]aninmals.Aninmal, numberOfContestants)
	for i := 0; i < numberOfContestants; {
		newAninmal := aninmals.Create()
		if !aninmals.IsDuplicate(newAninmal, uniqueAninmals) {
			uniqueAninmals[i] = newAninmal
			i++
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	go DisplayRace(ctx, uniqueAninmals)

	for i := range uniqueAninmals {
		wg.Add(1)
		go func(a *aninmals.Aninmal) {
			defer wg.Done()
			for {
				a.Race(rdb)
				if a.Progress >= raceLength {
					fmt.Printf("%s won the race!\n", a.Name)
					cancel() // Cancel the context to signal other goroutines
					return
				}
				time.Sleep(50 * time.Millisecond)

				select {
				case <-ctx.Done(): // Check if context is cancelled
					return
				default:
					// Continue racing
				}
			}
		}(&uniqueAninmals[i])
	}

	wg.Wait()
}
