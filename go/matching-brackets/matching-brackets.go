// Package brackets implements verifying that any and all pairs brackets are matched and nested correctly.
package brackets

// Bracket returns is there correct brackets positions.
func Bracket(in string) bool {

	var stack []rune

	for _, s := range in {
		if s == '(' || s == '[' || s == '{' {
			stack = append(stack, s)
			continue
		}
		if s == ')' || s == ']' || s == '}' {
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if !(pop == '(' && s == ')') && !(pop == '[' && s == ']') && !(pop == '{' && s == '}') {
				return false
			}
		}
	}

	return !(len(stack) > 0)
}
