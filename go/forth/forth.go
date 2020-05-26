// Package forth implements an evaluator for a very simple subset of Forth.
package forth

import (
	"errors"
	"strconv"
	"strings"
)

type evaluator struct {
	arg1, arg2 int
	operation  string
	result     int
	custom     map[string]string
}

// Forth returns evaluated result.
func Forth(in []string) ([]int, error) {
	var stack []int
	ev := newEval()
	command := in[len(in)-1]

	split := strings.Split(command, " ")
	for i, val := range split {
		arg, err := strconv.Atoi(val)
		if err != nil {
			if len(stack) < 2 {
				return nil, errors.New("invalid arguments")
			}

			ev.operation = val
			if err := ev.doOperation(); err != nil {
				return nil, err
			}

			stack = append(stack, ev.result)
			stack = stack[2:3]

			continue
		}

		stack = append(stack, arg)

		if i > 1 {
			continue
		}

		if i%2 == 0 {
			ev.arg1 = arg
			continue
		}

		ev.arg2 = arg
	}

	return stack, nil
}

func newEval() evaluator {
	return evaluator{
		custom: make(map[string]string),
	}
}

func (ev *evaluator) doOperation() error {
	switch ev.operation {
	case "+":
		ev.result = ev.arg1 + ev.arg2
	case "-":
		ev.result = ev.arg1 - ev.arg2
	case "*":
		ev.result = ev.arg1 * ev.arg2
	case "/":
		if ev.arg2 == 0 {
			return errors.New("divide by zero")
		}
		ev.result = ev.arg1 / ev.arg2
	default:
		return errors.New("invalid operation")
	}

	return nil
}
