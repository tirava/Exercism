// Package proverb implements generating the relevant proverb.
package proverb

import "fmt"

// Proverb returns relevant proverb from random words.
func Proverb(rhyme []string) []string {
	count := len(rhyme)
	proverb := make([]string, count)

	for i := 0; i < count-1; i++ {
		proverb[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
	}

	if count > 0 {
		proverb[count-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	}

	return proverb
}
