// Package accumulate implements different operations
// given a collection and an operation to perform on each element of the collection.
package accumulate

// Accumulate returns a new collection
// containing the result of applying the operation
// to each element of the input collection.
func Accumulate(in []string, f func(string) string) (out []string) {
	for _, word := range in {
		out = append(out, f(word))
	}
	return out
}
