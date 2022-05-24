package main

import (
	"log"
	"os"
	"errors"
	"github.com/Jeffail/gabs"
	"github.com/Murmuration-Labs/bitscreen"
	"github.com/ipfs/go-cid"
)

func getDealInfo() (cid.Cid, error) {
	var c cid.Cid
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
		[]string{"PayloadCID", "/"}, // retrieval deal
	} {
		c, err := cid.Parse(proposal.Search(path...).Data())
		// check only if found a valid CID
		if err == nil {
			return c, err
		}
	}
  return c, errors.New("No valid CID found.")
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
