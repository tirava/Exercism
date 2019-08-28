// Package robotname implements managing robot factory settings.
package robotname

// Robot base struct.
type Robot struct {
	name string
}

// Name returns robot name.
func (r *Robot) Name() (string, error) {

	return "", nil
}

// Reset resets robot name.
func (r *Robot) Reset() {

}
