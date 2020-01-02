// Package paasio implements reporting network IO statistics.
package paasio

import "io"

type paasCounter struct {
	readCount  int64
	writeCount int64
	writer     io.Writer
	reader     io.Reader
}

// NewWriteCounter returns new write struct.
func NewWriteCounter(w io.Writer) WriteCounter {

	return &paasCounter{writer: w}
}

// NewReadCounter returns new read struct.
func NewReadCounter(r io.Reader) ReadCounter {

	return &paasCounter{reader: r}
}

// NewReadWriteCounter returns new read and write struct.
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {

	return nil
}

func (pc *paasCounter) Write(p []byte) (n int, err error) {

	return pc.writer.Write(p)
}

func (pc *paasCounter) WriteCount() (n int64, nops int) {

	return 0, 0
}

func (pc *paasCounter) Read(p []byte) (n int, err error) {

	return pc.reader.Read(p)
}

func (pc *paasCounter) ReadCount() (n int64, nops int) {

	return 0, 0
}
