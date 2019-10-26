// Package flatten implements single flattened list with all values except nil/null.
package flatten

import (
	"reflect"
)

var out = make([]interface{}, 0)
var out1 = make([]interface{}, 0)
var nested int

// Flatten gets nested list and return a single flattened list.
func Flatten(in interface{}) []interface{} {

	var s reflect.Value
	//i := 0

	s = reflect.ValueOf(in)

	if in == nil {
		nested--
		return nil
	}

	switch reflect.TypeOf(in).Kind() {
	case reflect.Slice:
		//s = reflect.ValueOf(in)

		for i := 0; i < s.Len(); i++ {
			nested++
			Flatten(s.Index(i).Interface())
		}
	default:
		//fmt.Println("value:", s.Interface())
		if s.Interface() == nil {
			return nil
		}
		out = append(out, s.Interface())
	}
	nested--
	if nested < 0 {
		out1 = out
		out = make([]interface{}, 0)
		nested = 0
	}
	return out1
}
