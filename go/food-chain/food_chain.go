// Package foodchain implements.
package foodchain

var animal = []struct{ name, quote string }{
	{"", ""},
	{"fly", "I don't know why she swallowed the fly. Perhaps she'll die.\n"},
	{"spider", "It wriggled and jiggled and tickled inside her.\n"},
	{"bird", "How absurd to swallow a bird!\n"},
	{"cat", "Imagine that, to swallow a cat!\n"},
	{"dog", "What a hog, to swallow a dog!\n"},
	{"goat", "Just opened her throat and swallowed a goat!\n"},
	{"cow", "I don't know how she swallowed a cow!\n"},
	{"horse", "She's dead, of course!\n"}}

// Verse returns verse.
func Verse(n int) (r string) {
	r += "I know an old lady who swallowed a " + animal[n].name + ".\n"

	if n > 1 {
		r += animal[n].quote
	}

	if n < 8 {
		for i := n; i > 1; i-- {
			r += "She swallowed the " + animal[i].name + " to catch the " + animal[i-1].name
			if i == 3 { // bird to catch the spider
				r += " that wriggled and jiggled and tickled inside her.\n"
			} else {
				r += ".\n"
			}
		}

		r += animal[1].quote
	}

	return r[:len(r)-1]
}

// Verses returns verses.
func Verses(a, b int) (r string) {
	for i := a; i <= b; i++ {
		if i != 1 {
			r += "\n\n"
		}
		r += Verse(i)
	}
	return
}

// Song returns song.
func Song() string {
	return Verses(1, 8)
}
