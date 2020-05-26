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
	for _, val := range split {
		var validOperation bool

		arg, err := strconv.Atoi(val)
		if err != nil {
			//var validOperation bool

			if _, ok := ev.custom[val]; ok {
				validOperation = true
			}

			ev.operation = val

			if len(stack) < 2 {
				if !validOperation {
					return nil, errors.New("invalid arguments")
				}

				if len(stack) == 0 {
					return nil, errors.New("empty arguments")
				}

				ev.arg1 = stack[0]
			} else {
				if validOperation {
					ev.arg1 = stack[len(stack)-1]
				} else {
					ev.arg1, ev.arg2 = stack[0], stack[1]
				}
			}

			if err := ev.doOperation(); err != nil {
				return nil, err
			}

			//fmt.Println("stack before result:", stack, ev.arg1, ev.arg2, ev.operation)
			stack = append(stack, ev.result)
			//fmt.Println("stack before cut:", stack)
			if !validOperation {
				stack = stack[2:3]
				continue
			}

			//fmt.Println("stack:", stack, val)

			//stack = stack[2:3]
			//continue
		}
		if !validOperation {
			stack = append(stack, arg)
		}
	}

	return stack, nil
}

func newEval() evaluator {
	return evaluator{
		custom: map[string]string{
			"dup": "dup",
		},
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
	case "dup":
		ev.result = ev.arg1
	default:
		return errors.New("invalid operation")
	}

	return nil
}
