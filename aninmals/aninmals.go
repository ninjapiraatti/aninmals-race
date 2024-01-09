package aninmals

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis"
)

var (
	adjectives   = []string{"Loving", "Timid", "Furious", "Shiny", "Mechanical", "Pissed", "Cuddly"}
	aninmalNames = []string{"Treefloof", "Murder Mittens", "Patience Monkey", "Forest Gorgi", "Wizard Cow", "Formal Chikcen"}
)

type Aninmal struct {
	Name              string
	Progress          int
	Color             string
	ProgressFromRedis string
}

var colors = []string{
	"\033[31m", // Red
	"\033[32m", // Green
	"\033[33m", // Yellow
	"\033[34m", // Blue
	"\033[35m", // Magenta
	"\033[36m", // Cyan
}

func (a *Aninmal) Race(rdb *redis.Client) {
	status := rdb.Set(a.Name, a.Progress, 600*time.Second)
	if status.Err() != nil {
		fmt.Println("Redis error:", status)
	}
	val, err := rdb.Get(a.Name).Result()
	if err != nil {
		panic(err)
	}
	advance := rand.Intn(3)
	if a.Progress > 1 {
		advance -= 1
	}
	a.Progress += advance
	a.ProgressFromRedis = val
}

func Create() Aninmal {
	adjective := adjectives[rand.Intn(len(adjectives))]
	name := aninmalNames[rand.Intn(len(aninmalNames))]
	aninmal := Aninmal{
		Name:     adjective + " " + name,
		Progress: 0,
		Color:    colors[rand.Intn(len(colors))],
	}
	return aninmal
}

func IsDuplicate(aninmal Aninmal, aninmalList []Aninmal) bool {
	for _, a := range aninmalList {
		if aninmal.Name == a.Name {
			return true
		}
	}
	return false
}
