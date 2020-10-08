package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Murmuration-Labs/bitscreen"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/ipfs/go-cid"
)

func _handleErr(e error) error {
	log.Fatal(e)
	return e
}

func _buildCID() (cid.Cid, error) {
	return cid.Decode("bafzbeigai3eoy2ccc7ybwjfz5r3rdxqrinwi4rwytly24tdbh6yk7zslrm")
}

func _buildDeal() (storagemarket.MinerDeal, error) {
	j, err := os.Open("./test-deal.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Opened test-deal.json")
	defer j.Close()

	b, _ := ioutil.ReadAll(j)

	var d storagemarket.MinerDeal

	if e := json.Unmarshal(b, &d); e != nil {
		return d, _handleErr(e)
	}

	return d, nil
}

func main() {
	// Build test data
	d, e := _buildDeal()
	if e != nil {
		_handleErr(e)
	}

	bitscreen.ScreenDealProposal(d)
}
