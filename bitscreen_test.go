package bitscreen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/ipfs/go-cid"
	"gotest.tools/assert"
	"github.com/Jeffail/gabs"
)

func ScreenDealProposal(deal storagemarket.MinerDeal) bool {
	cid := deal.ProposalCid
	return BlockCidFromFile(cid)
}

func _handleErr(e error) error {
	log.Fatal(e)
	return e
}

func _buildDeal() (storagemarket.MinerDeal, error) {
	f := "./test_deal.json"
	j, err := os.Open(f)
	if err != nil {
		fmt.Println(err)
	}
	defer j.Close()

	b, _ := ioutil.ReadAll(j)

	var d storagemarket.MinerDeal

	if e := json.Unmarshal(b, &d); e != nil {
		return d, _handleErr(e)
	}

	return d, nil
}

func _maybeDeleteBitscreen() {
	p := GetPath()

	if FileExists(p) {
		e := os.RemoveAll(p)
		if e != nil {
			sigterm(e)
		}
	}
}

func _setCID(c cid.Cid) (cid.Cid, error) {
	p := GetPath()

	json := gabs.New()
	json.Array()
	json.ArrayAppend(c.String())

	return c, ioutil.WriteFile(p, []byte(json.String()+"\n"), os.ModePerm)
}

func _setEmptyCidList() (error) {
    p := GetPath()

    json := gabs.New()
    json.Array()

    return ioutil.WriteFile(p, []byte(json.String()+"\n"), os.ModePerm)
}

func TestScreenDealProposalExists(t *testing.T) {
	os.Setenv("BITSCREEN_FILENAME", "test/bitscreen")

	// Test CID exists in bitscreen, should return 1
	MaybeCreateBitscreen()
	d, e := _buildDeal()
	if e != nil {
		_handleErr(e)
	}
	_setCID(d.ProposalCid)

	result := ScreenDealProposal(d)
	assert.Equal(t, result, true)
	_maybeDeleteBitscreen()
}

func TestScreenDealProposalNotExists(t *testing.T) {
	os.Setenv("BITSCREEN_FILENAME", "test/bitscreen")

	// Test CID does not exist in bitscreen, should return 0
	d, e := _buildDeal()
	if e != nil {
		_handleErr(e)
	}

	_setEmptyCidList()

	result := ScreenDealProposal(d)
	assert.Equal(t, result, false)
}
