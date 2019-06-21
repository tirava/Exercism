// Package strain implements the keep and discard operation on collections.
package strain

type Ints []int
type Lists [][]int
type Strings []string

// Keep returns a new collection of Ints containing those elements where the predicate is true.
func (in Ints) Keep(f func(int) bool) (out Ints) {
	for _, val := range in {
		if f(val) {
			out = append(out, val)
		}
	}
	return out
}

// Discard returns a new collection of Ints containing those elements where the predicate is false.
func (in Ints) Discard(f func(int) bool) (out Ints) {
	for _, val := range in {
		if !f(val) {
			out = append(out, val)
		}
	}
	return out
}

// Keep returns a new collection of Lists containing those elements where the predicate is true.
func (in Lists) Keep(f func([]int) bool) (out Lists) {
	for _, val := range in {
		if f(val) {
			out = append(out, val)
		}
	}
	return out
}

// Keep returns a new collection of Strings containing those elements where the predicate is true.
func (in Strings) Keep(f func(string) bool) (out Strings) {
	for _, val := range in {
		if f(val) {
			out = append(out, val)
		}
	}
	return out
}

//func testVal(in interface{}, f func(interface{}) bool) (out interface{}) {
//
//	switch reflect.TypeOf(in).Kind() {
//	case reflect.Slice:
//		s := reflect.ValueOf(in)
//		for i := 0; i < s.Len(); i++ {
//			fmt.Println(s.Index(i).Interface())
//			//if f(val) {
//				//out = append(out, val)
//			//}
//		}
//	}
//	return out
//}
