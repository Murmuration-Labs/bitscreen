package main

import (
	"log"
	"os"
	"github.com/Jeffail/gabs"
	"github.com/Murmuration-Labs/bitscreen"
	"github.com/ipfs/go-cid"
)

func getDealInfo() (cid.Cid, error) {
	var c cid.Cid

	proposal, err := gabs.ParseJSONBuffer(os.Stdin)
	if err != nil {
		log.Fatalf("Unable to parse proposal JSON: %s", err)
	}

	for _, path := range [][]string{
		[]string{"Proposal", "PieceCID", "/"},
		[]string{"Proposal", "Label"},
		[]string{"Ref", "Root", "/"},
		[]string{"PayloadCID", "/"},
	} {
		c, err := cid.Parse(proposal.Search(path...).Data())
		// check only if found a valid CID
		if err == nil {
			return c, err
		}
	}
    return c, err
}

func main() {
    cid, err := getDealInfo()
    if err != nil {
        os.Exit(1)
    }

    if bitscreen.IsLoadFromFileEnabled() {
        if bitscreen.BlockCidFromFile(cid) {
            os.Exit(1)
        }
    } else {
        if bitscreen.BlockCidFromProcess(cid) {
            os.Exit(1)
        }
    }
}
