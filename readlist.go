package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func sigterm(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

func FindCid(cid string, p string) (bool, error) {
	f, err := os.Open(p)
	sigterm(err)
	defer f.Close()

	s := bufio.NewScanner(f)
	sigterm(s.Err())

	found := false

	for {
		b := s.Scan()
		if b {
			if s.Text() == cid {
				found = true
			}
		}

		if !b {
			break
		}
	}

	return found, nil
}

// /*
//  * Example implementation
//  */
// func main() {
// 	cid := "hello"
// 	p := "/tmp/dat" // Path to list of CIDs
// 	found, err := FindCid(cid, p)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(found)
// }
