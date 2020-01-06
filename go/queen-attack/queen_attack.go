// Package queenattack implements given the position of two queens on a chess board,
// indicate whether or not they are positioned so that they can attack each other.
package queenattack

import "errors"

// CanQueenAttack returns attack answer or error.
func CanQueenAttack(w, b string) (bool, error) {

	if len(w) > 2 || len(b) > 2 || w == b || len(w) < 2 || len(b) < 2 {
		return false, errors.New("invalid queen position length")
	}

	xw, yw := int(w[0]-0x60), int(w[1]-0x30)
	xb, yb := int(b[0]-0x60), int(b[1]-0x30)

	if xw < 1 || yw < 1 || xb < 1 || yb < 1 || xw > 8 || yw > 8 || xb > 8 || yb > 8 {
		return false, errors.New("invalid queen position coords")
	}

	if xw == xb || yw == yb {
		return true, nil
	}

	absW, absB := (yw-1)*8 + xw, (yb-1)*8 + xb

	if (absW - absB) % 7 == 0 || (absW - absB) % 9 == 0 {
		return true, nil
	}

	return false, nil
}
