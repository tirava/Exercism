// Package flatten implements single flattened list with all values except nil/null.
package flatten

// Flatten gets nested list and return a single flattened list.
func Flatten(input interface{}) []interface{} {
	flat := make([]interface{}, 0)
	for _, elem := range input.([]interface{}) {
		if inner, ok := elem.([]interface{}); ok {
			flat = append(flat, Flatten(inner)...)
		} else if elem != nil {
			flat = append(flat, elem)
		}
	}
	return flat
}
