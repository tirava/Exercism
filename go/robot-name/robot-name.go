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

const (
	maxNumNames = 26 * 26 * 10 * 10 * 10 // 676000
	letters     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits      = "0123456789"
)

// Name returns robot name.
func (r *Robot) Name() (string, error) {

	// comment for benchmark
	if len(robots) >= maxNumNames {
		return "", fmt.Errorf("no names for new robots: %d names already", len(robots))
	}

	if ok := robots[r.name]; ok != nil {
		return r.name, nil
	}

	b := make([]byte, 5)

	for {

		for i := range b {
			if i < 2 {
				b[i] = letters[rand.Int63()%int64(len(letters))] // Int63 much faster Intn
				continue
			}
			b[i] = digits[rand.Int63()%int64(len(digits))]
		}

		r.name = string(b)

		if _, ok := robots[r.name]; !ok {
			break
		}

		// uncomment for benchmark
		//if len(robots) >= maxNumNames {
		//	for i := 0; i < 6; i++ { // increase if error "reissued" occurred
		//		r.name += string(digits[rand.Int63()%int64(len(digits))])
		//	}
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
