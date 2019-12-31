// Package beer implements Recite the lyrics to that beloved classic,
// that field-trip favorite: 99 Bottles of Beer on the Wall.
package beer

import (
	"errors"
	"strconv"
	"strings"
)

// n1 = 99-1 or 'No more', $n2 =  99-1 or 'no more', s = '' or 's'
// n1-1 = 98-1, one = one or it
const (
	verse0991 = "$n1 bottle$s of beer on the wall, $n2 bottle$s of beer.\n"
	verse01   = "Take $one down and pass it around,"
	verse02   = "Go to the store and buy some more,"
	verse0993 = " $n1-1 bottle$s of beer on the wall.\n"
)

// Verse returns one verse.
func Verse(numVerse int) (string, error) {
	if numVerse < 0 || numVerse > 99 {
		return "", errors.New("invalid verse number")
	}

	v1 := verse0991
	v2 := verse01
	if numVerse == 0 {
		v2 = verse02
	}
	v3 := verse0993

	switch {
	case numVerse > 2:
		v3 = strings.Replace(v3, "$s", "s", -1)
		fallthrough
	case numVerse > 1:
		v1 = strings.Replace(v1, "$s", "s", -1)
		v2 = strings.Replace(v2, "$one", "one", 1)
		v3 = strings.Replace(v3, "$n1-1", strconv.Itoa(numVerse-1), 1)
		v3 = strings.Replace(v3, "$s", "", -1)
		fallthrough
	case numVerse > 0:
		v1 = strings.Replace(v1, "$n1", strconv.Itoa(numVerse), 1)
		v1 = strings.Replace(v1, "$n2", strconv.Itoa(numVerse), 1)
		fallthrough
	case numVerse == 1:
		v1 = strings.Replace(v1, "$s", "", -1)
		v2 = strings.Replace(v2, "$one", "it", 1)
		v3 = strings.Replace(v3, "$n1-1", "no more", 1)
		v3 = strings.Replace(v3, "$s", "s", -1)
	default: // == 0
		v1 = strings.Replace(v1, "$n1", "No more", 1)
		v1 = strings.Replace(v1, "$n2", "no more", 1)
		v1 = strings.Replace(v1, "$s", "s", -1)
		v3 = strings.Replace(v3, "$n1-1", "99", 1)
		v3 = strings.Replace(v3, "$s", "s", -1)
	}

	return v1 + v2 + v3, nil
}

// Verses returns verses from upper to lower.
func Verses(upper, lower int) (string, error) {
	if upper < lower {
		return "", errors.New("start < stop")
	}
	out := ""
	for i := upper; i >= lower; i-- {
		v, err := Verse(i)
		if err != nil {
			return "", errors.New("invalid verses numbers")
		}
		out += v + "\n"
	}
	return out, nil
}

// Song returns all beer song.
func Song() string {
	song, err := Verses(99, 0)
	if err != nil {
		return "invalid call of verses"
	}
	return song
}
