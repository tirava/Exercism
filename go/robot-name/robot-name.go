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

var robots = make(map[string]*Robot)

const maxNumNames = 676000 // 26 * 26 * 10 * 10 * 10

// Name returns robot name.
func (r *Robot) Name() (string, error) {

	// comment for benchmark
	if len(robots) >= maxNumNames {
		return "", fmt.Errorf("no names for new robots: %d names already", len(robots))
	}

	if ok := robots[r.name]; ok != nil {
		return r.name, nil
	}

	for {

		r1 := rune(rand.Intn('Z'-'A'+1)) + 'A'
		r2 := rune(rand.Intn('Z'-'A'+1)) + 'A'
		r3 := rune(rand.Intn('9'-'0'+1)) + '0'
		r4 := rune(rand.Intn('9'-'0'+1)) + '0'
		r5 := rune(rand.Intn('9'-'0'+1)) + '0'

		r.name = string(r1) + string(r2) + string(r3) + string(r4) + string(r5)

		if _, ok := robots[r.name]; !ok {
			break
		}

		// uncomment for benchmark
		//if len(robots) >= maxNumNames {
		//	r.name += string(rune(rand.Intn('9'-'0'+1)) + '0')
		//	r.name += string(rune(rand.Intn('9'-'0'+1)) + '0')
		//	r.name += string(rune(rand.Intn('9'-'0'+1)) + '0')
		//	r.name += string(rune(rand.Intn('9'-'0'+1)) + '0')
		//	r.name += string(rune(rand.Intn('9'-'0'+1)) + '0')
		//	break
		//}

	}

	robots[r.name] = r

	return r.name, nil
}

// Reset resets robot name.
func (r *Robot) Reset() {
	if r.name != "" {
		robots[r.name] = nil
	}
}
