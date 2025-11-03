package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(c byte) byte {
	switch {
	case 'A' <= c && c <= 'Z':
		return 'A' + (c-'A'+13)%26
	case 'a' <= c && c <= 'z':
		return 'a' + (c-'a'+13)%26
	default:
		return c
	}
}

func (r rot13Reader) Read(p []byte) (int, error) {
	for {
		n, err := r.r.Read(p)
		for i := 0; i < n; i++ {
			p[i] = rot13(p[i])
		}
		return n, err
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
