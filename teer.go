// teer.go - wrapper around io.ReadWrite(Close)r that copies data to itself.
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to teer, using the creative
// commons "cc0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package teer

import (
	"errors"
	"io"
)

type TeeReadWriter struct {
	rw io.ReadWriter
	tr io.Reader
}

func New(rw io.ReadWriter) *TeeReadWriter {
	return &TeeReadWriter{
		rw: rw,
		tr: io.TeeReader(rw, rw),
	}
}

func (t *TeeReadWriter) Read(b []byte) (int, error) {
	return t.tr.Read(b)
}

func (t *TeeReadWriter) Write(b []byte) (int, error) {
	return t.rw.Write(b)
}

func (t *TeeReadWriter) Close() error {
	switch t.rw.(type) {
	case io.ReadWriteCloser:
		return t.rw.(io.ReadWriteCloser).Close()
	}
	return errors.New("unable to close: parent does not implement Close()")
}
