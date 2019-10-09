// Package erratum implements various kinds of error handling and resource management.
package erratum

// Use opens a resource,
// calls Frob(input) on the result resource
// and then closes that resource.
func Use(o ResourceOpener, input string) (err error) {

	var res Resource

	res, err = o()
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		res, err = o()
	}
	defer res.Close()

	defer func() {
		if e := recover(); e != nil {
			if fe, ok := e.(FrobError); ok {
				res.Defrob(fe.defrobTag)
			}
			err = e.(error)
		}
	}()

	res.Frob(input)

	return err
}
