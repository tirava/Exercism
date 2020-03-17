// Package bowling implements the scoring and status saving of the game of bowling.
package bowling

import "errors"

// Game implements the type that store the status of the game
type Game struct {
	frameN          int
	score           int
	scoreMultiplier []int
	startPins       int
	shoot           int
	fitBall         int
	finished        bool
}

// NewGame create a new game of bowling.
func NewGame() *Game {
	return &Game{frameN: 1, startPins: 10, scoreMultiplier: []int{1, 1, 1}}
}

// Roll implements the effects on the score of a roll of a ball.
func (g *Game) Roll(pins int) error {
	if g.frameN == 11 {
		return errors.New("cannot roll an ended game")
	}
	g.startPins -= pins
	g.shoot++
	if g.startPins < 0 || pins > 10 || pins < 0 {
		return errors.New("invalid roll")
	}
	// different handle last shoot
	if g.frameN == 10 && g.startPins == 0 {
		g.startPins = 10
		g.fitBall = 1
	}
	g.score += pins * g.scoreMultiplier[0]
	g.scoreMultiplier = append(g.scoreMultiplier, 1)
	g.scoreMultiplier = g.scoreMultiplier[1:]
	// strike and spare
	if g.startPins == 0 {
		g.scoreMultiplier[0]++
		g.scoreMultiplier[1] += 2 - g.shoot
	}

	if g.shoot == 2+g.fitBall || (g.shoot == 1 && g.startPins == 0) {
		g.shoot = 0
		g.frameN++
		g.startPins = 10
	}

	if g.frameN == 11 {
		g.finished = true
	}

	return nil
}

// Score retrieve the current game score.
func (g *Game) Score() (int, error) {
	if g.finished != true {
		return 0, errors.New("incomplete game")
	}
	return g.score, nil
}
