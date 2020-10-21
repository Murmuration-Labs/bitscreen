package bitscreen

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-fil-markets/storagemarket"
)

// BitScreen holds information about the bitscreen file path
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

// FileExists checks whether a directory or file exists
func FileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// MaybeCreateBitscreen creates the ./murmuration/bitscreen file
// if the path does not already exist
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

// ScreenDealProposal compares a CID identified in the deal with
// the list of CIDs in the bitscreen
func ScreenDealProposal(deal storagemarket.MinerDeal) int {
	cid := deal.ProposalCid
	fmt.Printf("CID: %s\r\n", cid.String())
	return ScreenCID(cid)
}

// ScreenCID checks for a CID in ./murmuration/bitscreen
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
				fmt.Printf("Deals for CID %s are not welcome.\r\n", cid.String())
				return 1
			}
		} else {
			break
		}
	}
	return 0
}
