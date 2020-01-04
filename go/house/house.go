// Package house implements reciting the nursery rhyme
// 'This is the House that Jack Built'.
package house

var firstLines = [...]string{
	"This is the house that Jack built.",
	"This is the malt",
	"This is the rat",
	"This is the cat",
	"This is the dog",
	"This is the cow with the crumpled horn",
	"This is the maiden all forlorn",
	"This is the man all tattered and torn",
	"This is the priest all shaven and shorn",
	"This is the rooster that crowed in the morn",
	"This is the farmer sowing his corn",
	"This is the horse and the hound and the horn",
}

var otherLines = [...]string{
	"that lay in the house that Jack built.",
	"that ate the malt",
	"that killed the rat",
	"that worried the cat",
	"that tossed the dog",
	"that milked the cow with the crumpled horn",
	"that kissed the maiden all forlorn",
	"that married the man all tattered and torn",
	"that woke the priest all shaven and shorn",
	"that kept the rooster that crowed in the morn",
	"that belonged to the farmer sowing his corn",
}

// Verse returns concrete verse of the song.
func Verse(n int) string {

	others := ""

	for i := n - 2; i < 0; i-- {
		others += "\n" + otherLines[i]
	}

	return firstLines[n-1] + others
}

// Song returns all song.
func Song() string {
	result := ""
	for i := 1; i < 13; i++ {
		result += Verse(i) + "\n"
	}
	return result
}

// builder
