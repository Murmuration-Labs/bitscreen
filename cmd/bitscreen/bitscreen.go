package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/Murmuration-Labs/bitscreen"
	"github.com/ipfs/go-cid"
)

func _log(str string) {
	f, err := os.OpenFile("/tmp/bitscreen_go.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	log.Println(str)
}

func _logErr(str string) {
	f, err := os.OpenFile("/tmp/bitscreen_go.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	log.Fatalf(str)
}

func extractSelector(rawSelector string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9|]+")
	if err != nil {
		log.Fatal(err)
	}

	selector, _ := base64.StdEncoding.DecodeString(rawSelector)
	strSelector := reg.ReplaceAllString(string(selector), "")

	i := strings.Index(strSelector, "Raledepth")

	if i > -1 {
		return "", errors.New("Not a retrieval deal.")
	}

	splitIndex := strings.Index(strSelector, "|")

	if splitIndex > -1 {
		strSelector = strSelector[:splitIndex]
	}

	reg, err = regexp.Compile("afbf.")
	if err != nil {
		log.Fatal(err)
	}

	strSelector = reg.ReplaceAllString(strSelector, "/")

	strSelector = strings.Replace(strSelector, "Raldnoneb", "", -1)

	strSelector = strings.Trim(strSelector, "/")
	strSelector = strings.Trim(strSelector, "a")
	strSelector = strings.Trim(strSelector, "/Hash")

	return strSelector, nil
}

func extractFileCID(selector string, payloadCid cid.Cid) (cid.Cid, error) {
	var c cid.Cid
	var err error

	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	numbers := re.FindAllString(selector, -1)

	index := ""

	if len(numbers) > 0 {
		index = numbers[len(numbers)-1]

		newSelector := strings.Trim(selector, "/"+index)

		cmd := exec.Command("lotus", "client", "ls", "--data-selector", newSelector, payloadCid.String())

		var out bytes.Buffer
		var stdErr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stdErr

		err := cmd.Run()

		if err != nil {
			_log(out.String())
			_log(stdErr.String())
			log.Fatal(err)
		}

		//TODO: parse response to get CID (the nth line from the response where n is index variable)
	}

	_log(index)

	return c, err
}

func getDealInfo() ([]cid.Cid, error) {
	var cids []cid.Cid

	proposal, err := gabs.ParseJSONBuffer(os.Stdin)
	if err != nil {
		_logErr(fmt.Sprintf("Unable to parse proposal JSON: %s", err))
	}

	for _, path := range [][]string{
		[]string{"Ref", "Root", "/"},          // storage deal
		[]string{"Proposal", "PieceCID", "/"}, // storage deal
		[]string{"PayloadCID", "/"},           // retrieval deal
	} {
		c, err := cid.Parse(proposal.Search(path...).Data())
		// check only if found a valid CID
		if err == nil {
			cids = append(cids, c)
		}
	}

	rawSelector := proposal.Search("Selector", "Raw").Data()

	if rawSelector != nil {
		strSelector, err := extractSelector(rawSelector.(string))

		if err == nil {
			_, _ = extractFileCID(strSelector, cids[0])
			//TODO: add CID to cids array
		}
	}

	var cidErr error

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
		fmt.Println("Checking " + cid.String())
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
