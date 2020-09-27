package paasio

import "io"

// NewWriteCounter returns a new WriteCounter
func NewWriteCounter(w io.Writer) WriteCounter {
	return WriteCounter{io.Writer: w}
}

func NewReadCounter(r io.Reader) ReadCounter {
	return ReadCounter{io.Reader: r}
}

func NewReadWriteCounter(rw readWriter) {

}
