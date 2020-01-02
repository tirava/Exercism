// Package paasio implements reporting network IO statistics.
package paasio

import (
	"io"
	"sync"
)

type paasCounter struct {
	sync.RWMutex
	readCount      int64
	writeCount     int64
	nopsReadCount  int
	nopsWriteCount int
	writer         io.Writer
	reader         io.Reader
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
	return &paasCounter{
		writer: rw,
		reader: rw,
	}
}

func (pc *paasCounter) Write(p []byte) (int, error) {
	n, err := pc.writer.Write(p)
	pc.Lock()
	pc.writeCount += int64(n)
	pc.nopsWriteCount++
	pc.Unlock()
	return n, err
}

func (pc *paasCounter) WriteCount() (n int64, nops int) {
	pc.Lock()
	defer pc.Unlock()
	return pc.writeCount, pc.nopsWriteCount
}

func (pc *paasCounter) Read(p []byte) (int, error) {
	n, err := pc.reader.Read(p)
	pc.Lock()
	pc.readCount += int64(n)
	pc.nopsReadCount++
	pc.Unlock()
	return n, err
}

func (pc *paasCounter) ReadCount() (n int64, nops int) {
	pc.Lock()
	defer pc.Unlock()
	return pc.readCount, pc.nopsReadCount
}
