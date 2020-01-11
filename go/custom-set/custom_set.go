// Package stringset implements a custom data structure of some type (set).
package stringset

import (
	"fmt"
	"strings"
)

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
	s = removeDuplicates(s)
	elems := make([]string, len(s))
	copy(elems, s)
	return Set{elems: elems}
}

// String returns string view of the set.
func (s Set) String() string {
	sb := strings.Builder{}
	for _, e := range s.elems {
		sb.WriteString(fmt.Sprintf("\"%s\", ", e))
	}
	return fmt.Sprintf("{%s}", strings.TrimRight(sb.String(), ", "))
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

func removeDuplicates(s []string) []string {
	seen := make(map[string]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}
