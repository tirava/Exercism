// Package school implements graduating school students.
package school

// Grade is the base grade struct.
type Grade struct {
	grade    int
	students []string
}

// School is the base school type.
type School struct{}

// New returns new school.
func New() *School {
	return &School{}
}

// Add adds grade and students.
func (s *School) Add(string, int) {

}

// Grade grades students.
func (s *School) Grade(int) []string {
	return nil
}

// Enrollment enroll students.
func (s *School) Enrollment() []Grade {
	return nil
}
