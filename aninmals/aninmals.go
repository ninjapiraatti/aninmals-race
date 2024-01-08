package aninmals

import (
	"math/rand"
)

type Aninmal struct {
	Name string
}

var (
	adjectives   = []string{"Loving", "Timid", "Furious", "Shiny", "Mechanical", "Pissed", "Cuddly"}
	aninmalNames = []string{"Treefloof", "Murder Mittens", "Patience Monkey", "Forest Gorgi", "Wizard Cow", "Formal Chikcen"}
)

func Create() Aninmal {
	adjective := adjectives[rand.Intn(len(adjectives))]
	name := aninmalNames[rand.Intn(len(aninmalNames))]
	aninmal := Aninmal{
		Name: adjective + " " + name,
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
