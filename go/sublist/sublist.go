// Package sublist implements determine
// if the first list is contained within the second list
// and vise versa.
package sublist

// Relation is the base comparision type.
type Relation string

// Sublist returns result of relation of two lists.
func Sublist(list1, list2 []int) Relation {

	if len(list1) <= len(list2) {
		return Relation(list(list1, list2, "sub"))
	}
	return Relation(list(list2, list1, "super"))
}

func list(a, b []int, s string) string {
	for i := 0; i < len(b); i++ {
		if i+len(a) > len(b) {
			return "unequal"
		}
		bb := b[i : i+len(a)]
		if equal(a, bb) {
			if i == 0 && len(b) == len(a) {
				break
			} else {
				return s + "list"
			}
		}
	}
	return "equal"
}

func equal(a, b []int) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
