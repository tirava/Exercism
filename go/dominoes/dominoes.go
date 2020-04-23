// Package dominoes implements making a chain of dominoes.
package dominoes

// Domino base type.
type Domino [2]int

// MakeChain computes a way to order a given set of dominoes.
func MakeChain(input []Domino) (chain []Domino, ok bool) {
	if len(input) == 0 || (len(input) == 1 && input[0][0] == input[0][1]) {
		return input, true
	}

	return nil, false
}
