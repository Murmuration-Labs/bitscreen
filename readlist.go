package listtools

import (
	"bufio"
	"log"
	"os"
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"
)

func sigterm(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

func FindCid(cid cid.Cid, p string) (bool, error) {
	f, err := os.Open(p)
	sigterm(err)
	defer f.Close()

	s := bufio.NewScanner(f)
	sigterm(s.Err())

	for {
		b := s.Scan()
		if b {
			if s.Text() == cid.String() {
				return true, xerrors.New("cid found in list.")
			}

		if !b {
			break
		}
	}

	return false, nil
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
