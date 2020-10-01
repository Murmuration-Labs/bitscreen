# BitScreen

BitScreen is a tool to compare a CID with a list of CIDs. When a CID is found, an error is thrown.



## Usage

Example implementation:
```go
package main

import (
	"fmt"
	"log"

	"github.com/ipfs/go-cid"
	"github.com/Murmuration-Labs/bitscreen"
)

func main() {
	// CID to detect
	c, err := cid.Decode("bafzbeigai3eoy2ccc7ybwjfz5r3rdxqrinwi4rwytly24tdbh6yk7zslrm")
	if err != nil {
		log.Fatal(err)
	}

	// Screens CID
	found, err := bitscreen.Screen(c)

	// Handles found CID
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(found)
}
```
