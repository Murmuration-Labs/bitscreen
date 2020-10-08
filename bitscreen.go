package bitscreen

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/specs-actors/actors/builtin/market"
)

// For now, BitScreen holds information about the bitscreen file path
// This may expand in the future with additional functionality
type BitScreen struct {
	d  string
	fn string
	p  string
}

// Creates BitScreen struct
func getBitscreen() BitScreen {
	d := ".murmuration"
	fn := "bitscreen"
	b := BitScreen{
		d,
		fn,
		fmt.Sprintf("./%s/%s", d, fn),
	}

	return b
}

// Handle errors
func sigterm(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

// Checks whether directory or file exists
func FileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// If the ./murmuration/bitscreen file path does not exist, create it
func MaybeCreateBitscreen() BitScreen {
	b := getBitscreen()

	if !FileExists(b.d) {
		os.MkdirAll(b.d, 0777)
	}

	if !FileExists(b.p) {
		err := ioutil.WriteFile(b.p, []byte(""), 0777)
		sigterm(err)
	}

	return b
}

func ScreenClientDealProposal(deal market.ClientDealProposal) int {
	cid := deal.Proposal.PieceCID
	return ScreenCID(cid)
}

// Checks for a CID in ./murmuration/bitscreen
// If content should be filtered, returns 0
// If content should not be filtered, returns 1
func ScreenCID(cid cid.Cid) int {
	b := MaybeCreateBitscreen()
	f, err := os.OpenFile(b.p, os.O_RDONLY, 0777)
	sigterm(err)
	defer f.Close()

	s := bufio.NewScanner(f)
	sigterm(s.Err())

	for {
		b := s.Scan()
		if b {
			if s.Text() == cid.String() {
				return 0
			}
		} else {
			break
		}
	}

	return 1
}
