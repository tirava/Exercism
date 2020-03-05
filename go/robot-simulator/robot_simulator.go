// Package robot implements robot simulator.
package robot

// String is stringer for direction.
func (d Dir) String() string {
	return string(d)
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

// Action base type.
type Action string

// StartRobot turns robot on.
func StartRobot(cmd chan Command, act chan Action) {
	for c := range cmd {
		switch c {
		case 'R':
			act <- "Right"
		case 'L':
			act <- "Left"
		case 'A':
			act <- "Advance"
		}
	}

	println("S2R close")
	close(act)
}

// Room inits room.
func Room(extent Rect, robot Step2Robot, act chan Action, rep chan Step2Robot) {
	s2r := Step2Robot{}
	Step1Robot.Dir = robot.Dir
	Step1Robot.X, Step1Robot.Y = int(robot.Easting), int(robot.Northing)

	for a := range act {
		switch a {
		case "Right":
			Right()
		case "Left":
			Left()
		case "Advance":
			Advance()
		}
	}

	s2r.Dir = Step1Robot.Dir
	s2r.Pos = Pos{RU(Step1Robot.X), RU(Step1Robot.Y)}
	println("Sent report")
	rep <- s2r
}
