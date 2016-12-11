// Rot-13 Reader
// @author Prabhash Rathore

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (myRot13Reader rot13Reader) Read(b []byte) (int, error) {
	inputByteLength := len(b)
	a := make([]byte, inputByteLength)
	
	count, err := myRot13Reader.r.Read(a)
	
	if err != nil {
		return 0, err
	}
	for i:= 0; i < count; i++ {
		// add Rot-13 logic here
		c := a[i]
		if c >= 65  && c <= 90 {
			c = c - 65
			c = c + 13
			c = c % 26
			c = c + 65
		} else if c >= 97 && c <= 122 {
			c = c - 97
			c = c + 13
			c = c % 26
			c = c + 97
		}
		
		b[i] = c
	}
	
	return count, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

