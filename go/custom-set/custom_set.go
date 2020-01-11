// Package stringset implements a custom data structure of some type (set).
package stringset

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements.  For example {"a", "b"}.
// Format the empty set as {}.

// Set is the base set struct.
type Set struct {
	elems []string
}

// New returns new set.
func New() Set {
	return Set{elems: make([]string, 0)}
}

// NewFromSlice returns set from slice.
func NewFromSlice(s []string) Set {

	return Set{}
}

// String returns string view of the set.
func (s Set) String() string {

	return ""
}

// IsEmpty returns empty flag.
func (s Set) IsEmpty() bool {

	return false
}

// Has return has element flag.
func (s Set) Has(string) bool {

	return false
}

// Subset returns sub elements in set.
func Subset(s1, s2 Set) bool {

	return false
}

// Disjoint returns flag is sets disjoint.
func Disjoint(s1, s2 Set) bool {

	return false
}

// Equal returns flag is sets equal.
func Equal(s1, s2 Set) bool {

	return false
}

// Add adds string into set.
func (s Set) Add(string) {

}

// Intersection returns set of inter sets.
func Intersection(s1, s2 Set) Set {

	return Set{}
}

// Difference returns diff of the sets.
func Difference(s1, s2 Set) Set {

	return Set{}
}

// Union returns union of the sets.
func Union(s1, s2 Set) Set {

	return Set{}
}
