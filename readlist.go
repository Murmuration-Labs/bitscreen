package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

type ListReader struct {
	p string         // File path to read
	f *os.File       // File data
	s *bufio.Scanner // Scanner
}

func (lr *ListReader) Next() (string, bool) {
	b := lr.s.Scan()

	if b {
		return lr.s.Text(), b
	}

	return "", b
}

func NewListReader(p string) *ListReader {
	f, err := os.Open(p)
	check(err)

	s := bufio.NewScanner(f)
	check(s.Err())

	lr := ListReader{
		p: p,
		f: f,
		s: s,
	}
	return &lr
}

// Example of how to use NewListReader
func main() {
	p := "/tmp/dat"
	lr := NewListReader(p)

	// Loop through each line and print
	for {
		d, b := lr.Next()
		fmt.Println(d)

		if !b {
			break
		}
	}

	// Important step - don't forget this!
	// There's probably a better way to handle it.
	defer lr.f.Close()
}
