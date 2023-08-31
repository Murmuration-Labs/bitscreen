package main

import (
	"fmt"
	"os"
	"testing"

	"gotest.tools/assert"
	"github.com/ipfs/go-cid"

// 	"github.com/Jeffail/gabs"
)

func _contains(s []cid.Cid, e string) bool {
    for _, a := range s {
        if a.String() == e {
            return true
        }
    }
    return false
}

func TestGetDealInfoStorageDeal(t *testing.T) {
    f := "../../testdata/storage_proposal.json"
	j, err := os.Open(f)
	if err != nil {
		fmt.Println(err)
	}
	defer j.Close()

    oldStdin := os.Stdin
    defer func() { os.Stdin = oldStdin }()

    os.Stdin = j

    c, err := getDealInfo()

	assert.Equal(t, _contains(c, "baga6ea4seaqnxqw4jbd7qytukwtwyj634tk53ic7bvam7c77tqiqsb5zcxsvikq"), true)
	assert.Equal(t, _contains(c, "bafykbzacea4laqix45psjv43pi436sx3ghybf6zvyvcxozvuqxo23dxbqiho2"), true)
	assert.Equal(t, _contains(c, "bafyreihktt5tqkzivkfxnq3d2q2kdqast4aymgrsixfqbatywdrxkf6umm"), false)
}

func TestGetDealInfoRetrievalDeal(t *testing.T) {
    f := "../../testdata/retrieval_proposal.json"
	j, err := os.Open(f)
	if err != nil {
		fmt.Println(err)
	}
	defer j.Close()

    oldStdin := os.Stdin
    defer func() { os.Stdin = oldStdin }()

    os.Stdin = j

    c, err := getDealInfo()

    assert.Equal(t, _contains(c, "bafykbzacea4laqix45psjv43pi436sx3ghybf6zvyvcxozvuqxo23dxbqiho2"), true)
}