package circular

import (
	"errors"
)

type Buffer struct {
	buf   []byte
	read  int // next position for read
	write int // next position for write
}

var (
	ErrReadEmpty  = errors.New("Error reading empty circular buffer")
	ErrFullBuffer = errors.New("Error writing into full buffer")
)

// NewBuffer creates a new circular buffer of size specified and returns a pointer
func NewBuffer(size int) *Buffer {
	return &Buffer{buf: make([]byte, size)}
}

// ReadByte reads the next byte in the buffer and returns the result or an error
func (b *Buffer) ReadByte() (byte, error) {
	if b.write == b.read { // buffer is empty
		return 0, ErrReadEmpty
	}
	readByte := b.buf[b.read%len(b.buf)]
	b.read++
	return readByte, nil
}

// WriteByte write a new byte into the buffer and outputs an error
func (b *Buffer) WriteByte(c byte) error {
	if (b.write)%len(b.buf) == b.read && b.buf[b.write%len(b.buf)] != 0 {
		return ErrFullBuffer
	}
	b.buf[b.write%len(b.buf)] = c
	b.write++
	return nil
}

// Overwrite overwrite the oldest byte with the new byte
func (b *Buffer) Overwrite(c byte) {
	if (b.write)%len(b.buf) == b.read && b.buf[b.write%len(b.buf)] != 0 {
		b.read++
	}
	b.buf[b.write%len(b.buf)] = c
	b.write++

}

// Reset put buffer in an empty state
func (b *Buffer) Reset() {
	for i := range b.buf {
		b.buf[i] = 0
	}
	b.read = 0
	b.write = 0
}
