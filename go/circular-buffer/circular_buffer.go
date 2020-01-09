// Package circular implements circular buffer, cyclic buffer or ring buffer.
package circular

import (
	"errors"
)

// Buffer is the base buffer type.
type Buffer struct {
	begin  int
	end    int
	size   int
	free   int
	buffer []byte
}

// NewBuffer return new buffer.
func NewBuffer(size int) *Buffer {
	return &Buffer{
		size:   size,
		free:   size,
		buffer: make([]byte, size),
	}
}

// ReadByte returns read byte or error.
func (b *Buffer) ReadByte() (byte, error) {
	if b.free == b.size {
		return 0, errors.New("buffer is empty")
	}
	read := b.buffer[b.begin]
	b.buffer[b.begin] = 0
	b.free++
	b.begin++
	if b.begin > b.size-1 {
		b.begin = 0
	}
	return read, nil
}

// WriteByte writes byte into buffer.
func (b *Buffer) WriteByte(c byte) error {
	if b.free == 0 {
		return errors.New("buffer is full")
	}
	b.free--
	b.buffer[b.end] = c
	b.end++
	if b.end > b.size-1 {
		b.end = 0
	}
	return nil
}

// Overwrite overwrites byte in buffer.
func (b *Buffer) Overwrite(c byte) {
	if b.free > 0 {
		_ = b.WriteByte(c)
		return
	}
	b.begin++
	if b.begin > b.size-1 {
		b.begin = 0
	}
	b.buffer[b.end] = c
	b.end++
	if b.end > b.size-1 {
		b.end = 0
	}
}

// Reset puts buffer in empty state.
func (b *Buffer) Reset() {
	b.begin = 0
	b.end = 0
	b.free = b.size
	b.buffer = nil
	b.buffer = make([]byte, b.size)
}
