// Package letter implements concurrency counting letters in some languages.
package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

const maxLettersInAlphabet = 48

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency concurrently
func ConcurrentFrequency(s []string) FreqMap {

	ls := len(s)                             // faster & less alloc
	m := make(FreqMap, maxLettersInAlphabet) // result
	wg := &sync.WaitGroup{}
	ch := make(chan FreqMap, ls) // channel for merging maps

	// goroutine for merging maps
	wg.Add(1)
	go func(ch chan FreqMap) {
		for i := 0; i < ls; i++ {
			fm := <-ch
			for k, v := range fm {
				m[k] += v
			}
		}
		wg.Done()
	}(ch)

	// goroutines for calc letters
	for i := range s {
		wg.Add(1)
		go func(text *string) {
			ch <- Frequency(*text)
			wg.Done()
		}(&s[i])
	}

	wg.Wait()

	return m
}
