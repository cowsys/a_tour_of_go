package main

import (
	"io"
	"os"
	"regexp"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)

	var s string
	for i := 0; i < n; i++ {
		s = string(b[i])

		if m, _ := regexp.MatchString("[A-Ma-m]", s); m {
			b[i] = b[i] + 13
		} else if m, _ = regexp.MatchString("[N-Zn-z]", s); m {
			b[i] = b[i] - 13
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penxrq gur pbqr!")
	r := rot13Reader{s}

	io.Copy(os.Stdout, &r)
}
