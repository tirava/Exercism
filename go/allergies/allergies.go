// Package allergies implements determining whether or not they're allergic to a given item,
// and their full list of allergies.
package allergies

var words = map[uint]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

// Allergies returns slice of allergens
func Allergies(code uint) (als []string) {

	for i := uint(1); i <= 128; i <<= 1 {
		if code&i > 0 {
			als = append(als, words[i])
		}
	}

	return
}

// AllergicTo returns slice of concrete allergens
func AllergicTo(code uint, al string) bool {

	for _, als := range Allergies(code) {
		if als == al {
			return true
		}
	}

	return false
}
