// Package react implements a basic reactive system.
package react

type react struct {
	ic []inputCell
	*computeCell
}

type inputCell struct {
	cell
	cc  *computeCell
	ics *[]inputCell
}

type computeCell struct {
	cell
	cancel
}

type cancel struct {
	callback func(int)
}

type cell struct {
	value    int
	compute1 func(int) int
	compute2 func(int, int) int
}

// New returns new react.
func New() Reactor {
	c := cell{
		compute1: func(i int) int {
			return i
		},
		compute2: func(i, j int) int {
			return 0
		},
	}

	cc := computeCell{
		cell: c,
	}

	//ic := inputCell{
	//	cell: c,
	//	cc:   &cc,
	//}

	r := &react{
		computeCell: &cc,
		ic:          make([]inputCell, 0, 2),
	}

	//r.ic = append(r.ic, ic)

	return r
}

// CreateInput creates an input cell.
func (r *react) CreateInput(value int) InputCell {
	//c := cell{
	//	compute1: func(i int) int {
	//		return i
	//	},
	//	compute2: func(i, j int) int {
	//		return 0
	//	},
	//}

	ic := inputCell{
		cell: cell{value: value},
		cc:   r.computeCell,
		ics:  &r.ic,
	}

	r.ic = append(r.ic, ic)
	//r.ic[len(r.ic)-1].SetValue(value)
	//r.ic[len(r.ic)-1].value = value
	//fmt.Printf("cretae: %p %d\n", &r.ic[len(r.ic)-1], r.ic[len(r.ic)-1].value)

	return &r.ic[len(r.ic)-1]
}

// CreateCompute1 creates a compute cell 1.
func (r *react) CreateCompute1(cell Cell, calc func(int) int) ComputeCell {
	r.computeCell.cell.compute1 = calc
	//r.ic[0].SetValue(cell.Value())
	r.ic[0].value = cell.Value()
	r.computeCell.cell.value = r.computeCell.cell.compute1(cell.Value())

	return r.computeCell
}

// CreateCompute1 creates a compute cell 2.
func (r *react) CreateCompute2(c1 Cell, c2 Cell, calc func(int, int) int) ComputeCell {
	r.computeCell.cell.compute2 = calc
	//r.inputCell.SetValue(c1.Value(), c2.Value())
	r.computeCell.cell.value = r.computeCell.cell.compute2(c1.Value(), c2.Value())

	//fmt.Println(c1.Value(), c2.Value(), "->", r.computeCell.cell.value)

	return r.computeCell
}

// SetValue sets the value of the cell.
func (ic *inputCell) SetValue(value int) {
	ic.cell.value = value
	if len(*ic.ics) == 1 {
		ic.cc.cell.value = ic.cc.cell.compute1(value)

		return
	}

	for i := range *ic.ics {
		//fmt.Printf("I: %d, addresses: %p %d %p %d\n", i, &(*ic.ics)[i], (*ic.ics)[i].value, ic, ic.cell.value)
		if &(*ic.ics)[i] != ic {
			ic.cc.cell.value = ic.cc.cell.compute2(value, (*ic.ics)[i].value)
			//	//fmt.Println("addresses:", &c, ic)
		}
		//} else {
		//	//fmt.Println("---------------", &c, ic)
		//}
	}

	//ic.cc.cell.value = ic.cc.cell.compute2(value, ic.ics[1])
	//ic.cc.cell.value = ic.cc.cell.compute2(value)
	//fmt.Println(value, "->", ic.cc.cell.value, *ic.ics)
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
