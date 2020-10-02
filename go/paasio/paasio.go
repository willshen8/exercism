package paasio

import (
	"io"
	"sync"
)

type counter struct {
	n      int64
	nops   int
	reader io.Reader
	writer io.Writer
	mutex  sync.Mutex
}

// ReadCount returns the number of bytes read and the number of operations
func (c *counter) ReadCount() (n int64, nops int) {
	return c.n, c.nops
}

// WriteCount returns the number of written bytes and the number of operations
func (c *counter) WriteCount() (n int64, nops int) {
	return c.n, c.nops
}

func (c *counter) Read(input []byte) (int, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	n, err := c.reader.Read(input)
	if err != nil {
		return 0, err
	}
	c.n += int64(n)
	c.nops++
	return n, nil
}

func (c *counter) Write(input []byte) (int, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	n, err := c.writer.Write(input)
	if err != nil {
		return 0, err
	}
	c.n += int64(n)
	c.nops++
	return n, nil
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &counter{reader: r}
}

// NewWriteCounter returns a new WriteCounter
func NewWriteCounter(w io.Writer) WriteCounter {
	return &counter{writer: w}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &counter{reader: rw, writer: rw}
}
