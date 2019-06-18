// Package scale implements generating the musical scale
// starting with the tonic and following the specified interval pattern
package scale

import "strings"

// Scale generates scales
func Scale(tonic, interval string) (s []string) {

	var scale []string

	switch tonic {
	case "C", "a", "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#":
		scale = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		scale = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
	}

	tonic = strings.Title(tonic)
	for i, note := range scale {
		if tonic == note {
			s = append(s, scale[i:]...)
			s = append(s, scale[0:i]...)
		}
	}

	if interval != "" {
		i := 0
		var s1 []string
		for _, note := range interval {
			s1 = append(s1, s[i]) // s[] may be nil? - check this in the feature
			switch note {
			case 'm':
				i++
			case 'M':
				i += 2
			case 'A':
				i += 3
			}
		}
		s = s1
	}

	return s
}
