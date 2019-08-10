// Package letter implements concurrency counting letters in some languages.
package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

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

	type FreqMapCon struct {
		sync.RWMutex
		FreqMap
	}

	m := FreqMapCon{}
	m.FreqMap = make(FreqMap)
	wg := &sync.WaitGroup{}

	for _, text := range s {
		wg.Add(1)
		go func(text string) {
			m.Lock()
			for _, r := range text {
				m.FreqMap[r]++
			}
			m.Unlock()
			wg.Done()
		}(text)
	}

	wg.Wait()

	return m.FreqMap
}
