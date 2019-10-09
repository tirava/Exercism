// Package erratum implements various kinds of error handling and resource management.
package erratum

// Use opens a resource,
// calls Frob(input) on the result resource
// and then closes that resource.
func Use(o ResourceOpener, input string) (err error) {

	var res Resource

	defer func() {
		if e := recover(); e != nil {
			if fe, ok := e.(FrobError); ok {
				res.Defrob(fe.defrobTag)
				err = fe.inner
			} else {
				err = e.(error)
			}
			_ = res.Close()
		}
	}()

	for {
		res, err = o()
		if err != nil {
			if _, ok := err.(TransientError); !ok {
				return err
			}
		} else {
			break
		}
	}

	res.Frob(input)

	return res.Close()
}
