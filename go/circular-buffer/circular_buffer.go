// Package circular implements circular buffer, cyclic buffer or ring buffer.
package circular

import "errors"

type Buffer struct {
	pointer int
	size    int
	free    int
	buffer  []byte
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
	read := b.buffer[b.pointer]
	b.free++
	b.pointer--
	if b.pointer < 0 {
		b.pointer = b.size - 1
	}
	return read, nil
}

// WriteByte writes byte into buffer.
func (b *Buffer) WriteByte(c byte) error {
	if b.free == 0 {
		return errors.New("buffer is full")
	}
	b.free--
	b.pointer++
	if b.pointer > b.size-1 {
		b.pointer = 0
	}
	b.buffer[b.pointer] = c
	return nil
}

// Overwrite overwrites byte in buffer.
func (b *Buffer) Overwrite(c byte) {

}

// Reset puts buffer in empty state.
func (b *Buffer) Reset() {

}
