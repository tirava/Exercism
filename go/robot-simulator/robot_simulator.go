// Package robot implements robot simulator.
package robot

// String is stringer for direction.
func (d Dir) String() string {
	return "111"
}

// global directions.
const (
	N Dir = iota
	S
	W
	E
)

// Left stepping.
func Left() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = W
	case W:
		Step1Robot.Dir = S
	case S:
		Step1Robot.Dir = E
	case E:
		Step1Robot.Dir = N
	}
}

// Right stepping.
func Right() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = E
	case E:
		Step1Robot.Dir = S
	case S:
		Step1Robot.Dir = W
	case W:
		Step1Robot.Dir = N
	}
}

// Advance for stepping.
func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case S:
		Step1Robot.Y--
	case E:
		Step1Robot.X++
	case W:
		Step1Robot.X--
	}
}
