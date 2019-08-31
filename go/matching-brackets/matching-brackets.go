// Package brackets implements verifying that any and all pairs brackets are matched and nested correctly.
package brackets

// Bracket returns is there correct brackets positions.
func Bracket(in string) bool {

	var stack []rune

	for _, s := range in {
		//println(string(s))
		if s == '(' || s == '[' || s == '{' {
			//println("append stack:", string(s))
			stack = append(stack, s)
			continue
		}
		if s == ')' || s == ']' || s == '}' {
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			//println("pop stack:", string(pop))
			stack = stack[0:len(stack)-1]
			if (pop == '(' && s == ')') || (pop == '[' && s == ']') || (pop == '{' && s == '}') {
				//println(string(pop), string(s))
				println(pop, s)
				continue
			}
		}
	}

	//fmt.Println(stack)
	if len(stack) > 0 {
		return false
	}

	return true
}