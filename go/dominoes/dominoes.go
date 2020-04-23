// Package dominoes implements making a chain of dominoes.
package dominoes

import "fmt"

// Domino base type.
type Domino [2]int

// MakeChain computes a way to order a given set of dominoes.
func MakeChain(input []Domino) (chain []Domino, ok bool) {
	if len(input) == 0 || (len(input) == 1 && input[0][0] == input[0][1]) {
		return input, true
	}

	chain = make([]Domino, 1, len(input))
	chain[0] = input[0]

	in := make([]Domino, len(input))
	copy(in, input)

	var j int

	for i := 1; i < len(input); i++ {
		var found bool
		var revert Domino

		if chain[j][1] == input[i][0] {
			revert = input[i]
			found = true
		} else if chain[j][1] == input[i][1] {
			revert = input[i]
			revert[0], revert[1] = revert[1], revert[0]
			found = true
		}

		if found {
			chain = append(chain, revert)
			in = append(in[:i], in[i+1:]...)
			j++
		}

		if i == len(input)-1 && len(in) > 1 {
			i = 0
		}
	}

	fmt.Println("chain:", chain)
	fmt.Println("in:", in)

	if chain[0][0] == chain[len(chain)-1][1] {
		return chain, true
	}

	return nil, false
}
