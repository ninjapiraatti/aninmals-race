package aninmals

import (
	"math/rand"
)

var (
	adjectives   = []string{"Loving", "Timid", "Furious", "Shiny", "Mechanical", "Pissed", "Cuddly"}
	aninmalNames = []string{"Treefloof", "Murder Mittens", "Patience Monkey", "Forest Gorgi", "Wizard Cow", "Formal Chikcen"}
)

type Aninmal struct {
	Name     string
	Progress int
	Color    string
}

func generateColor() string {
	return "\033[31m" // Red
}

func (a *Aninmal) Race() {
	advance := rand.Intn(3)
	a.Progress += advance
}

func Create() Aninmal {
	adjective := adjectives[rand.Intn(len(adjectives))]
	name := aninmalNames[rand.Intn(len(aninmalNames))]
	aninmal := Aninmal{
		Name:     adjective + " " + name,
		Progress: 0,
		Color:    generateColor(),
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
