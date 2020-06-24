// Package react implements a basic reactive system.
package react

type react struct {
	inputCell
	computeCell
}

type inputCell struct {
	cell
}

type computeCell struct {
	cell
	cancel
}

type cancel struct {
}

type cell struct {
}

// New returns new react.
func New() Reactor {
	return react{}
}

// CreateInput creates an input cell.
func (r react) CreateInput(int) InputCell {
	return r.inputCell
}

// CreateCompute1 creates a compute cell 1.
func (r react) CreateCompute1(Cell, func(int) int) ComputeCell {
	return r.computeCell
}

// CreateCompute1 creates a compute cell 2.
func (r react) CreateCompute2(Cell, Cell, func(int, int) int) ComputeCell {
	return r.computeCell
}

// SetValue sets the value of the cell.
func (ic inputCell) SetValue(int) {

}

// Value returns cell value.
func (c cell) Value() int {
	return 0
}

// AddCallback adds a callback.
func (cc computeCell) AddCallback(func(int)) Canceler {
	return cc.cancel
}

// Cancel removes callbacks.
func (c cancel) Cancel() {

}
