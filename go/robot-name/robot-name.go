// Package robotname implements managing robot factory settings.
package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Robot base struct.
type Robot struct {
	name string
}

var robots = make(map[string]bool)

const maxNumNames = 26 * 26 * 10 * 10 * 10 // 676000

// Name returns robot name.
func (r *Robot) Name() (string, error) {

	if r.name != "" {
		return r.name, nil
	}

	if len(robots) >= maxNumNames {
		return "", fmt.Errorf("no names for new robots: %d names already", len(robots))
	}

	r.name = newName()
	for robots[r.name] {
		r.name = newName()
	}

	robots[r.name] = true

	return r.name, nil
}

func newName() string {
	r1 := string(rand.Intn(26) + 'A')
	r2 := string(rand.Intn(26) + 'A')
	num := rand.Intn(1000)
	return fmt.Sprintf("%s%s%03d", r1, r2, num)
}

// Reset resets robot name.
func (r *Robot) Reset() {
	r.name = ""
}
