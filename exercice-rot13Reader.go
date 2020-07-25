package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if err != nil {
		return
	}
	for i := 0; i < n; i++ {
		switch {
		case 'a' <= p[i] && p[i] <= 'z':
			p[i] = 'a' + (p[i]-'a'+13)%26
		case 'A' <= p[i] && p[i] <= 'Z':
			p[i] = 'A' + (p[i]-'A'+13)%26
		}
	}
	return
}

func exerciceRot13Reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	_, _ = io.Copy(os.Stdout, &r)
}
