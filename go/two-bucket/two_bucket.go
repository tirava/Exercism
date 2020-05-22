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

	if sizeBucketOne <= 0 || sizeBucketTwo <= 0 || goalAmount <= 0 ||
		startBucket != "one" && startBucket != "two" {
		return "", 0, 0, errors.New("invalid data")
	}

	if sizeBucketTwo/sizeBucketOne > 1 && sizeBucketOne != 1 {
		return "", 0, 0, errors.New("not solving")
	}

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

	for {
		// pure to big while big not full
		for b2.now != b2.size {
			// pure to big
			g1, g2 := b1.pure(&b2)
			if g1 {
				return "one", steps, b2.now, nil
			} else if g2 {
				return "two", steps, b1.now, nil
			}

			// fill small
			if b1.fill() {
				return "one", steps, b2.now, nil
			}
		}

		b2.empty()
	}
}

func bigFirst(b1, b2 bucket) (
	goalBucket string, numSteps, otherBucketLevel int, e error) {

	for {
		// fill big
		if b2.fill() {
			return "two", steps, b1.now, nil
		}

		// pure to small + empty small while big not empty
		for b2.now != 0 {
			g2, g1 := b2.pure(&b1)
			if g1 {
				return "one", steps, b2.now, nil
			} else if g2 {
				return "two", steps, b1.now, nil
			}

			b1.empty()
		}
	}
}

func (b *bucket) fill() bool {
	if b.now > 0 {
		return false
	}

	steps++
	b.now = b.size

	return b.now == b.goal
}

func (b *bucket) empty() {
	if b.size == b.now {
		steps++
		b.now = 0
	}
}

func (b *bucket) pure(bx *bucket) (g1, g2 bool) {
	if b.now == 0 {
		return false, false
	}

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
