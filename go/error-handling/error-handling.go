// Package erratum implements various kinds of error handling and resource management.
package erratum

// Use opens a resource,
// calls Frob(input) on the result resource
// and then closes that resource.
func Use(o ResourceOpener, input string) error {

	defer func() {
		err := recover()
		if err != nil {
			println("rrrrrrrrrrrrrrrrrrrr")
			//res.Frob(input)
		}
	}()

	res, err := o()
	if _, ok := err.(TransientError); !ok {
		return err
	}

	println("ffffffffffffffff")

	res.Frob(input)

	res.Close()

	return nil
}
