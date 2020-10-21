package bitscreen

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-fil-markets/storagemarket"
)

// GetPath returns the filepath to the bitscreen
func GetPath() string {
	r := ".murmuration"
	fn, exists := os.LookupEnv("BITSCREEN_FILENAME")
	if !exists {
		fn = "bitscreen"
	}

	return filepath.Join(r, fn)
}

// MaybeCreateBitscreen generates instance of BitScreen struct
// and if needed, creates a bitscreen file
func MaybeCreateBitscreen() bool {
	p := GetPath()

	if !FileExists(p) {
		dir, _ := filepath.Split(p)
		os.MkdirAll(dir, os.ModePerm)
		err := ioutil.WriteFile(p, []byte(""), os.ModePerm)

		if err != nil {
			log.Fatal(err)
			return false
		}
	}

	return true
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

// ScreenDealProposal compares a CID identified in the deal with
// the list of CIDs in the bitscreen
func ScreenDealProposal(deal storagemarket.MinerDeal) int {
	cid := deal.ProposalCid
	return ScreenCID(cid)
}

// ScreenCID checks for a CID in ./murmuration/bitscreen
// If content should be filtered, returns 0
// If content should not be filtered, returns 1
func ScreenCID(cid cid.Cid) int {
	MaybeCreateBitscreen()
	p := GetPath()
	f, err := os.OpenFile(p, os.O_RDONLY, os.ModePerm)
	if err != nil {
		sigterm(err)
	}
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
