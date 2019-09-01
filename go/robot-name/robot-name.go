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

	// comment for benchmark
	if len(robots) >= maxNumNames {
		return "", fmt.Errorf("no names for new robots: %d names already", len(robots))
	}

	if robots[r.name] {
		return r.name, nil
	}

	for ok := true; ok; {
		r.name = r.newName()
		_, ok = robots[r.name]

		// uncomment for benchmark
		//if len(robots) >= maxNumNames {
		//	num := rand.Intn(1000000) // increase x10 and %0xd below if error "reissued" occurred
		//	r.name += fmt.Sprintf("%06d", num)
		//	break
		//}
	}

	robots[r.name] = true

	return r.name, nil
}

func (r *Robot) newName() string {
	r1 := string(rand.Intn(26) + 'A')
	r2 := string(rand.Intn(26) + 'A')
	num := rand.Intn(1000)
	return fmt.Sprintf("%s%s%03d", r1, r2, num)
}

// Reset resets robot name.
func (r *Robot) Reset() {
	if r.name != "" {
		robots[r.name] = false
	}
}
