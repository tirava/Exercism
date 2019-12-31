// Package beer implements Recite the lyrics to that beloved classic,
// that field-trip favorite: 99 Bottles of Beer on the Wall.
package beer

import (
	"errors"
	"fmt"
	"strings"
)

// Verse returns one verse.
func Verse(numVerse int) (string, error) {
	if numVerse < 0 || numVerse > 99 {
		return "", errors.New("invalid verse number")
	}

	switch numVerse {
	case 0:
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	case 1:
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	case 2:
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
	default:
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", numVerse, numVerse, numVerse-1), nil
	}
}

// Verses returns verses from upper to lower.
func Verses(upper, lower int) (string, error) {
	if upper > 100 || lower < 0 || lower > upper {
		return "", errors.New("start < stop")
	}
	sb := strings.Builder{}
	for i := upper; i >= lower; i-- {
		v, err := Verse(i)
		if err != nil {
			return "", errors.New("invalid verses numbers")
		}
		sb.WriteString(v)
		sb.WriteString("\n")
	}
	return sb.String(), nil
}

// Song returns all beer song.
func Song() string {
	song, _ := Verses(99, 0)
	return song
}
