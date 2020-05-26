// Package forth implements an evaluator for a very simple subset of Forth.
package forth

type evaluator struct {
	arg1, arg2 int
	operation  string
	result     int
	custom     map[string]string
}

// Forth returns evaluated result.
func Forth([]string) ([]int, error) {

	return nil, nil
}

func newEval() evaluator {
	return evaluator{
		custom: make(map[string]string),
	}
}

func (ev evaluator) doOperation() error {

	return nil
}
