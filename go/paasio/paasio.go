// Package paasio implements reporting network IO statistics.
package paasio

import "io"

type paasCounter struct {
	error
}

// NewWriteCounter returns new write struct.
func NewWriteCounter(w io.Writer) WriteCounter {

	return &paasCounter{}
}

// NewReadCounter returns new read struct.
func NewReadCounter(r io.Reader) ReadCounter {

	return nil
}

// NewReadWriteCounter returns new read and write struct.
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {

	return nil
}

func (w *paasCounter) Write(p []byte) (n int, err error) {

	return w.Write(p)
}

func (w *paasCounter) WriteCount() (n int64, nops int) {

	return 0, 0
}
