// Package kindergarten implements determine which plants each child
// in the kindergarten class is responsible for.
package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

var plantsLetters = map[rune]string{
	'G': "grass",
	'C': "clover",
	'R': "radishes",
	'V': "violets",
}

//var indexes = map[string]int{
//	"Alice":   1,
//	"Bob":     2,
//	"Charlie": 3,
//	"David":   4,
//	"Eve":     5,
//	"Fred":    6,
//	"Ginny":   7,
//	"Harriet": 8,
//	"Ileana":  9,
//	"Joseph":  10,
//	"Kincaid": 11,
//	"Larry":   12,
//}

// Garden is the base type.
type Garden struct {
	indexes map[string]int
	diagram []string
	//children []string
}

// NewGarden constructor returns new Garden.
func NewGarden(diagram string, children []string) (*Garden, error) {
	sort.Strings(children)
	g := &Garden{}

	g.diagram = strings.Split(diagram, "\n")
	if len(g.diagram) != 3 {
		return nil, errors.New("bad diagram")
	}

	if len(g.diagram[1])%2 != 0 || len(g.diagram[2])%2 != 0 {
		return nil, errors.New("bad len diagram")
	}

	if strings.ToLower(g.diagram[1]) == g.diagram[1] || strings.ToLower(g.diagram[2]) == g.diagram[2] {
		return nil, errors.New("bad caps diagram")
	}

	g.indexes = make(map[string]int, len(children))
	for i, c := range children {
		if c == "" {
			return nil, errors.New("no children")
		}
		if _, ok := g.indexes[c]; ok {
			return nil, errors.New("children exists")
		}
		g.indexes[c] = i + 1
	}
	return g, nil
}

// Plants returns plants for given child
func (g *Garden) Plants(child string) ([]string, bool) {
	gardens := make([]string, 4)

	index := (g.indexes[child] - 1) * 2
	if index < 0 {
		return nil, false
	}

	diag1 := g.diagram[1][index : index+2]
	diag2 := g.diagram[2][index : index+2]

	for i, d := range diag1 + diag2 {
		gardens[i] = plantsLetters[d]
		//fmt.Println(plantsLetters[d])
	}

	return gardens, true
}
