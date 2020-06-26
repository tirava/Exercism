// Package react implements a basic reactive system.
package react

type react struct {
	*inputCell
	*computeCell
}

type inputCell struct {
	cell
	cc *computeCell
}

type computeCell struct {
	cell
	cancel
}

type cancel struct {
	callback func(int)
}

type cell struct {
	value   int
	compute func(int) int
}

// New returns new react.
func New() Reactor {
	c := cell{
		compute: func(i int) int {
			return i
		},
	}

	cc := computeCell{
		cell: c,
	}

	ic := inputCell{
		cell: c,
		cc:   &cc,
	}

	return &react{
		computeCell: &cc,
		inputCell:   &ic,
	}
}

// CreateInput creates an input cell.
func (r *react) CreateInput(value int) InputCell {
	r.SetValue(value)

	return r.inputCell
}

// CreateCompute1 creates a compute cell 1.
func (r *react) CreateCompute1(cell Cell, calc func(int) int) ComputeCell {
	r.computeCell.cell.compute = calc
	r.inputCell.SetValue(cell.Value())

	return r.computeCell
}

// CreateCompute1 creates a compute cell 2.
func (r *react) CreateCompute2(Cell, Cell, func(int, int) int) ComputeCell {
	return r.computeCell
}

// SetValue sets the value of the cell.
func (ic *inputCell) SetValue(value int) {
	ic.cell.value = value
	ic.cc.cell.value = ic.cc.cell.compute(value)
}

// Value returns cell value.
func (c *cell) Value() int {
	return c.value
}

// AddCallback adds a callback.
func (cc *computeCell) AddCallback(cb func(int)) Canceler {
	//cc.callback = calc
	return &cc.cancel
}

// Cancel removes callbacks.
func (c *cancel) Cancel() {
	//c.callback = func(i int) int {
	//	return i
	//}
}
