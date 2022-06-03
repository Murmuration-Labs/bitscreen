package main

import (
	"log"
	"os"
	"errors"
	"github.com/Jeffail/gabs"
	"github.com/Murmuration-Labs/bitscreen"
	"github.com/ipfs/go-cid"
	"fmt"
)

func getDealInfo() ([]cid.Cid, error) {
	var cids []cid.Cid

	f, err := os.OpenFile("/tmp/bitscreen_go.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
        log.Fatalf("error opening file: %v", err)
    }
    defer f.Close()
    log.SetOutput(f);

	proposal, err := gabs.ParseJSONBuffer(os.Stdin)
    log.Println(proposal.String())
	if err != nil {
		log.Fatalf("Unable to parse proposal JSON: %s", err)
	}

	for _, path := range [][]string{
		[]string{"Ref", "Root", "/"}, // storage deal
		[]string{"Proposal", "PieceCID", "/"}, // storage deal
		[]string{"PayloadCID", "/"}, // retrieval deal
	} {
		c, err := cid.Parse(proposal.Search(path...).Data())
		// check only if found a valid CID
		if err == nil {
			cids = append(cids, c);
		}
	}

	var cidErr error;

    if len(cids) == 0 {
        cidErr = errors.New("No valid CID found.")
    }

    return cids, cidErr
}

func main() {
    cids, err := getDealInfo()
    if err != nil {
        os.Exit(1)
    }

    for _, cid := range cids {
        fmt.Println("Checking "+cid.String())
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
}
