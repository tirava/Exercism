// Package twelve implements output the lyrics to 'The Twelve Days of Christmas'.
package twelve

import "fmt"

var days = [...]string{
	"first", "second", "third", "fourth", "fifth", "sixth",
	"seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth",
}

var gifts = [...]string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

// Song returns all verses.
func Song() (song string) {
	for i := range days {
		song += Verse(i + 1)
		if i != len(days)-1 {
			song += "\n"
		}
	}
	return
}

// Verse returns n-th verse.
func Verse(in int) string {
	s := ""

	for i := 0; i < in; i++ {
		if i == 0 {
			s = gifts[i]
			if in > 1 {
				s = "and " + s
			}
		} else {
			s = gifts[i] + ", " + s
		}
	}

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", days[in-1], s)
}
