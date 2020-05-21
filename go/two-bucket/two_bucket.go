// Package twobucket demonstrates
// how to measure an exact number of liters
//by strategically transferring liters of fluid between the buckets.
package twobucket

import (
	"errors"
)

type bucket struct {
	size int
	now  int
	goal int
}

var steps int

// Solve solves the problem.
func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (
	goalBucket string, numSteps, otherBucketLevel int, e error) {

	steps = 0
	b1 := bucket{size: sizeBucketOne, goal: goalAmount}
	b2 := bucket{size: sizeBucketTwo, goal: goalAmount}

	if startBucket == "two" {
		return bigFirst(b1, b2)
	}

	return smallFirst(b1, b2)
}

func smallFirst(b1, b2 bucket) (
	goalBucket string, numSteps, otherBucketLevel int, e error) {

	// pure to big while big not full
	for b2.now != b2.size {
		// fill small
		if b1.fill() {
			return "one", steps, b2.now, nil
		}

		// pure to big
		g1, g2 := b1.pure(&b2)
		if g1 {
			return "one", steps, b2.now, nil
		} else if g2 {
			return "two", steps, b1.now, nil
		}
	}

	// empty big
	//b2.empty()

	// pure to big
	//b1.pure(&b2)
	//if g1 {
	//	return "one", steps, b2.now, nil
	//} else if g2 {
	//	return "two", steps, b1.now, nil
	//}

	return "", 0, 0, errors.New("todo smallFirst")
}

func bigFirst(b1, b2 bucket) (
	goalBucket string, numSteps, otherBucketLevel int, e error) {

	// fill big
	if b2.fill() {
		return "two", steps, b1.now, nil
	}

	// pure to small + empty small while big not empty
	for b2.now != 0 {
		g1, g2 := b2.pure(&b1)
		if g1 {
			return "one", steps, b2.now, nil
		} else if g2 {
			return "two", steps, b1.now, nil
		}

		b1.empty()
	}

	return "", 0, 0, errors.New("todo bigFirst")
}

func (b *bucket) fill() bool {
	steps++
	b.now = b.size

	return b.now == b.goal
}

func (b *bucket) empty() {
	steps++
	b.now = 0
}

func (b *bucket) pure(bx *bucket) (g1, g2 bool) {
	steps++
	if bx.now+b.now > bx.size {
		b.now -= bx.size - bx.now
		bx.now = bx.size
	} else {
		bx.now += b.now
		b.now = 0
	}

	return b.now == b.goal, bx.now == bx.goal
}
