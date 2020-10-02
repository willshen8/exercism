package paasio

import (
	"io"
	"sync"
)

type readCounter struct {
	n      int64
	nops   int
	reader io.Reader
	mutex  sync.Mutex
}

type writeCounter struct {
	n      int64
	nops   int
	writer io.Writer
	mutex  sync.Mutex
}

type readWriteCounter struct {
	readCounter
	writeCounter
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{reader: r}
}

// ReadCount returns the number of bytes read and the number of operations
func (r *readCounter) ReadCount() (n int64, nops int) {
	return r.n, r.nops
}

func (r *readCounter) Read(input []byte) (int, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	n, err := r.reader.Read(input)
	if err != nil {
		return 0, err
	}
	r.n += int64(n)
	r.nops++
	return n, nil
}

// NewWriteCounter returns a new WriteCounter
func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{writer: w}
}

// WriteCount returns the number of written bytes and the number of operations
func (w *writeCounter) WriteCount() (n int64, nops int) {
	return w.n, w.nops
}

func (w *writeCounter) Write(input []byte) (int, error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	n, err := w.writer.Write(input)
	if err != nil {
		return 0, err
	}
	w.n += int64(n)
	w.nops++
	return n, nil
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		readCounter{reader: rw},
		writeCounter{writer: rw},
	}
}
