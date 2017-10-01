package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13byte(c byte) byte {
	switch {
	case c >= 'a' && c <= 'z':
		c = 'a' + (c-'a'+13)%26
	case c >= 'A' && c <= 'Z':
		c = 'A' + (c-'A'+13)%26
	}
	return c
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)

	for i := 0; i < n; i++ {
		b[i] = rot13byte(b[i])
	}
	return len(b), err
}

func (r rot13Reader) Error() string {
	return fmt.Sprintf("Cannot read from: %v", r.r)
}

func main() {
	fmt.Printf("%c\n", rot13byte('a'))
	fmt.Printf("%c\n", rot13byte('z'))
	fmt.Printf("%c\n", rot13byte('A'))
	fmt.Printf("%c\n", rot13byte('Z'))
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
