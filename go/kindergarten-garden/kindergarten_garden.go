// Package kindergarten implements determine which plants each child
// in the kindergarten class is responsible for.
package kindergarten

var plantsLetters = map[string]string{
	"G": "grass",
	"C": "clover",
	"R": "radishes",
	"V": "violets",
}

var indexes = map[string]int{
	"Alice":   0,
	"Bob":     1,
	"Charlie": 2,
	"David":   3,
	"Eve":     4,
	"Fred":    5,
	"Ginny":   6,
	"Harriet": 7,
	"Ileana":  8,
	"Joseph":  9,
	"Kincaid": 10,
	"Larry":   11,
}

// Garden is the base type.
type Garden struct {
	diagram  string
	children []string
}

// NewGarden constructor returns new Garden.
func NewGarden(diagram string, children []string) (*Garden, error) {
	g := &Garden{
		diagram: diagram,
	}
	for _, c := range children {
		g.children = append(g.children, c)
	}
	return g, nil
}

// Plants returns plants for given child
func (g *Garden) Plants(child string) ([]string, bool) {

	return nil, false
}
