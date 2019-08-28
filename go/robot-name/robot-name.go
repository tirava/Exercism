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

var robots = make(map[string]*Robot, 0)

// Name returns robot name.
func (r *Robot) Name() (string, error) {

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
		//println("Collision:", r.name)
		//r.Reset()
		if len(robots) >= 676000 { //26*26*10*10*10 {
			return "", fmt.Errorf("end robots names: %d", len(robots))
		}
	}

	robots[r.name] = r
	//fmt.Println(r.name)

	return r.name, nil
}

// Reset resets robot name.
func (r *Robot) Reset() {
	//delete(robots, r.name)
	robots[r.name] = nil
}
