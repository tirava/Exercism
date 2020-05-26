// Package forth implements an evaluator for a very simple subset of Forth.
package forth

import (
	"errors"
	"strconv"
	"strings"
)

type evaluator struct {
	arg1, arg2  int
	operation   string
	result      int
	defOpers    map[string]string
	customOpers map[string]string
}

// Forth returns evaluated result.
func Forth(in []string) ([]int, error) {
	var (
		stack   []int
		command string
	)
	ev := newEval()

	command = strings.ToLower(in[len(in)-1])

	if len(in) > 1 {
		var arg1, arg2 string

		for i := 0; i < len(in)-1; i++ {
			in[i] = strings.ToLower(in[i])

			split := strings.Split(in[i], " ")
			if split[0] != ":" {
				return nil, errors.New("invalid custom prefix")
			}

			arg1 = split[1]
			_, err := strconv.Atoi(arg1)
			if err == nil {
				return nil, errors.New("invalid custom command")
			}

			arg2 = strings.Join(split[2:len(split)-1], " ")

			if _, ok := ev.customOpers[arg2]; ok {
				ev.customOpers[arg1] = ev.customOpers[arg2]
				continue
			}

			ev.customOpers[arg1] = arg2
		}

		//fmt.Println(ev.customOpers)
		for k, v := range ev.customOpers {
			command = strings.ReplaceAll(command, k, v)
		}
		//command = strings.ReplaceAll(command, arg1, arg2)
	}

	split := strings.Split(command, " ")
	for _, val := range split {
		var validOperation bool

		arg, err := strconv.Atoi(val)
		if err != nil {
			oper, ok := ev.defOpers[val]
			if ok {
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

			stack = append(stack, ev.result)

			if !validOperation {
				stack = stack[2:3]
				continue
			}

			switch oper {
			case "drop":
				stack = stack[:len(stack)-2]
			case "swap":
				if len(stack) < 3 {
					return nil, errors.New("too many arguments for swap")
				}

				stack = stack[:len(stack)-1]
				stack[len(stack)-1], stack[len(stack)-2] = stack[len(stack)-2], stack[len(stack)-1]
			case "over":
				if len(stack) < 3 {
					return nil, errors.New("too many arguments for over")
				}
				stack[len(stack)-1] = stack[len(stack)-3]
			}
		}

		if !validOperation {
			stack = append(stack, arg)
		}
	}

	return stack, nil
}

func newEval() evaluator {
	return evaluator{
		defOpers: map[string]string{
			"dup":  "dup",
			"drop": "drop",
			"swap": "swap",
			"over": "over",
		},
		customOpers: make(map[string]string),
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
	case "dup", "drop", "swap", "over":
		ev.result = ev.arg1
	default:
		return errors.New("invalid operation")
	}

	return nil
}
