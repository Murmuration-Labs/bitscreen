package main

import(
  "github.com/Murmuration-Labs/bitscreen"
)

func main() {
	// CID to detect
	c, err := cid.Decode("bafzbeigai3eoy2ccc7ybwjfz5r3rdxqrinwi4rwytly24tdbh6yk7zslrm")
	if err != nil {
    log.Fatal(err)
  }

	// Screens CID
	found, err := Screen(c)

	// Handles found CID
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(found)
}
