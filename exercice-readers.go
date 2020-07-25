package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (n int, err error) {
	for n = 0; n < len(b); n++ {
		b[n] = 'A'
	}
	return n, nil
}

func exerciceReaders() {
	reader.Validate(MyReader{})
}
