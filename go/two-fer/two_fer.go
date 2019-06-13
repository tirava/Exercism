// Package twofer implements simple "short for two for one"
package twofer

// ShareWith returns string with "you" (if name is blank) or with name
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
