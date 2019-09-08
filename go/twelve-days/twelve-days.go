// Package twelve implements output the lyrics to 'The Twelve Days of Christmas'.
package twelve

import (
	"fmt"
	"strings"
)

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
func Song() string {

	sb := strings.Builder{}

	for i := range days {
		sb.WriteString(Verse(i + 1))
		if i != len(days)-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

// Verse returns n-th verse.
func Verse(in int) string {

	sb := strings.Builder{}

	for i := in - 1; i >= 0; i-- {
		sb.WriteString(gifts[i])
		if i >= 1 {
			sb.WriteString(", ")
		}
		if i == 1 {
			sb.WriteString("and ")
		}
	}

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", days[in-1], sb.String())
}
