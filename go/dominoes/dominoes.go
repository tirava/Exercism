// Package dominoes implements making a chain of dominoes.
package dominoes

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

	for i := 1; i < len(in); i++ {
		var found bool
		var revert Domino

		if chain[j][1] == in[i][0] {
			revert = in[i]
			found = true
		} else if chain[j][1] == in[i][1] {
			revert = in[i]
			revert[0], revert[1] = revert[1], revert[0]
			found = true
		}

		if found {
			chain = append(chain, revert)
			in = append(in[:i], in[i+1:]...)
			j++
			i = 0
		}

		if len(in) > 1 {
			if !isDigitInDomino(in, chain[len(chain)-1][1]) {
				if chain[0][0] != chain[len(chain)-1][1] {
					return nil, false
				}

				chain = append(chain, chain[0])
				chain = chain[1:]
			}
		}
	}

	if chain[0][0] == chain[len(chain)-1][1] && len(in) == 1 {
		return chain, true
	}

	return nil, false
}

func isDigitInDomino(in []Domino, digit int) bool {
	for i := 1; i < len(in); i++ {
		if in[i][0] == digit || in[i][1] == digit {
			return true
		}
	}

	return false
}
