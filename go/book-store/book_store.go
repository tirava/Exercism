// Package bookstore implements trying and encourage
//more sales of different books from a popular 5 book series.
package bookstore

const max = 5
const price = 800

var discounts = [max]int{0, 5, 10, 20, 25}
var replacers = []struct {
	to   int
	from [2]int
}{{to: 3, from: [2]int{2, 4}}}

const bookLimit = 100000000

func min(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	if a > b {
		return a
	}
	return b
}

func counter(a []int) (m []int) {
	m = make([]int, max)
	for _, x := range a {
		m[x-1]++
	}
	return m
}

func floorMin(m *[]int) (zero, min int) {
	min = bookLimit
	for i := 0; i < max; i++ {
		if (*m)[i] < 1 {
			zero++
			continue
		}
		if (*m)[i] < min {
			min = (*m)[i]
		}
	}
	for i := 0; i < max; i++ {
		(*m)[i] -= min
	}
	return
}

func decOrder(count []int) (groups []int) {
	groups = make([]int, max)
	for i := 0; i < max; i++ {
		z, mn := floorMin(&count)
		if z == max {
			break
		}
		groups[max-z-1] = mn
	}
	// 2 groups of 4 better than a group of five and a four
	for _, rep := range replacers {
		add := min(groups[rep.from[0]], groups[rep.from[1]])
		groups[rep.from[0]] -= add
		groups[rep.from[1]] -= add
		groups[rep.to] += add * 2
	}
	return
}

// Cost calculates best books cost.
func Cost(books []int) (res int) {
	booksC := counter(books)
	groups := decOrder(booksC)
	for i, g := range groups {
		res += g * price * (i + 1) * (100 - discounts[i]) / 100
	}
	return res
}