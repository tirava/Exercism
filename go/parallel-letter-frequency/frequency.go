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

	m := make(FreqMap, maxLettersInAlphabet)
	fm := make([]FreqMap, len(s))
	wg := &sync.WaitGroup{}
	ch := make(chan FreqMap, len(s))

	//mu := &sync.RWMutex{}

	wg.Add(1)
	go func(ch chan FreqMap) {
		for i := 0; i < len(s); i++ {
			fm := <-ch
			for k, v := range fm {
				m[k] += v
			}
		}
		wg.Done()
	}(ch)

	for i := range s {
		fm[i] = make(FreqMap, maxLettersInAlphabet)
		wg.Add(1)
		go func(text *string, index int) {
			for _, r := range *text {
				fm[index][r]++
			}
			wg.Done()
			ch <- fm[index]
		}(&s[i], i)
	}

	wg.Wait()

	return m
}
