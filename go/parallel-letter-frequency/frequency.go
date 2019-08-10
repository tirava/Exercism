// Package letter implements concurrency counting letters in some languages.
package letter

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
	m := FreqMap{}
	for _, text := range s {
		mX := Frequency(text)
		for k, v := range mX {
			m[k] += v
		}
		//fmt.Println(m)
	}
	return m
}
