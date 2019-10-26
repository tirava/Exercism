// Package flatten implements single flattened list with all values except nil/null.
package flatten

import "reflect"

var out []interface{}
var nested int

// Flatten gets nested list and return a single flattened list.
func Flatten(in interface{}) []interface{} {

	if in == nil {
		nested--
		return nil
	}

	s := reflect.ValueOf(in)

	switch reflect.TypeOf(in).Kind() {
	case reflect.Slice:
		for i := 0; i < s.Len(); i++ {
			nested++
			Flatten(s.Index(i).Interface())
		}
	default:
		out = append(out, s.Interface())
	}

	nested--
	if nested < 0 {
		nested = 0
		outRet := out
		out = make([]interface{}, 0)
		return outRet
	}

	return nil
}

//func Flatten(input interface{}) []interface{} {
//	flat := make([]interface{}, 0)
//	if nested, ok := input.([]interface{}); ok {
//		for _, elem := range nested {
//			if inner, ok := elem.([]interface{}); ok {
//				flat = append(flat, Flatten(inner)...)
//			} else if elem != nil {
//				flat = append(flat, elem)
//			}
//		}
//	} else if input != nil {
//		flat = append(flat, input)
//	}
//	return flat
//}
