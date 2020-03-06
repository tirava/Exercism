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

		if Step1Robot.X < int(extent.Min.Easting) {
			Step1Robot.X = int(extent.Min.Easting)
		}
		if Step1Robot.Y < int(extent.Min.Northing) {
			Step1Robot.Y = int(extent.Min.Northing)
		}
		if Step1Robot.X > int(extent.Max.Easting) {
			Step1Robot.X = int(extent.Max.Easting)
		}
		if Step1Robot.Y > int(extent.Max.Northing) {
			Step1Robot.Y = int(extent.Max.Northing)
		}
	}

	s2r.Dir = Step1Robot.Dir
	s2r.Pos = Pos{RU(Step1Robot.X), RU(Step1Robot.Y)}

	rep <- s2r

	Step1Robot.X, Step1Robot.Y = 0, 0
	Step1Robot.Dir = N
}

// Action3 base type.
type Action3 struct {
	name   string
	action string
}

// StartRobot3 extends robots features.
func StartRobot3(name, script string, action chan Action3, log chan string) {
	if name == "" {
		log <- "A robot without a name"
		action <- Action3{name, "End"}
		return
	}

	for _, c := range script {
		switch c {
		case 'R':
			action <- Action3{name, "Right"}
		case 'L':
			action <- Action3{name, "Left"}
		case 'A':
			action <- Action3{name, "Advance"}
		}
	}

	action <- Action3{name, "End"}
}

// Room3 is extended room.
func Room3(extent Rect, robots []Step3Robot, action chan Action3, report chan []Step3Robot, log chan string) {
	s3r := make(map[string]*Step3Robot, len(robots))
	var isWall bool

	for i, r := range robots {
		if _, ok := s3r[r.Name]; ok {
			log <- "Duplicate robot names"
			report <- robots
			return
		}

		s3r[r.Name] = &robots[i]
	}

	count := len(robots)
	for a := range action {
		switch a.action {
		case "Right":
			Right3(s3r[a.name])
		case "Left":
			Left3(s3r[a.name])
		case "Advance":
			Advance3(s3r[a.name])
		case "End":
			count--
			if count == 0 {
				close(action)
			}
		}

		if s3r[a.name].Easting < extent.Min.Easting {
			s3r[a.name].Easting = extent.Min.Easting
			isWall = true
		}
		if s3r[a.name].Northing < extent.Min.Northing {
			s3r[a.name].Northing = extent.Min.Northing
			isWall = true
		}
		if s3r[a.name].Easting > extent.Max.Easting {
			s3r[a.name].Easting = extent.Max.Easting
			isWall = true
		}
		if s3r[a.name].Northing > extent.Max.Northing {
			s3r[a.name].Northing = extent.Max.Northing
			isWall = true
		}

		if isWall {
			isWall = false
			log <- "A robot attempting to advance into a wall"
		}
	}

	report <- robots
}

// Right3 stepping.
func Right3(robot *Step3Robot) {
	switch robot.Dir {
	case N:
		robot.Dir = E
	case E:
		robot.Dir = S
	case S:
		robot.Dir = W
	case W:
		robot.Dir = N
	}
}

// Left3 stepping.
func Left3(robot *Step3Robot) {
	switch robot.Dir {
	case N:
		robot.Dir = W
	case W:
		robot.Dir = S
	case S:
		robot.Dir = E
	case E:
		robot.Dir = N
	}
}

// Advance3 for stepping.
func Advance3(robot *Step3Robot) {
	switch robot.Dir {
	case N:
		robot.Northing++
	case S:
		robot.Northing--
	case E:
		robot.Easting++
	case W:
		robot.Easting--
	}
}
