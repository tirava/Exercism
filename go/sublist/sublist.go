// package sublists implements determine
// if the first list is contained within the second list
// and vise versa.
package sublist

type Relation string

// Sublist returns result of relation of two lists.
func Sublist(list1, list2 []int) Relation {
	len1, len2 := len(list1), len(list2)

	if len1 > 0 && len2 == 0 {
		return "superlist"
	}

	if len1 <= len2 {
		for i := 0; i < len2; i++ {
			if i+len1 > len2 {
				return "unequal"
			}
			b := list2[i : i+len1]
			if equal(list1, b) {
				if i == 0 && len2 == len1 {
					break
				} else {
					return "sublist"
				}
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
