package bitscreen

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"github.com/ipfs/go-cid"
	"github.com/pebbe/zmq4"
	"github.com/Jeffail/gabs"
)

type updaterResponse struct {
	reject int
	dealCid string
	cid string
	err string
}

/* Supported env vars */

//   BITSCREEN_FILENAME -- name of file containing the CIDs to block, defaults to `bitscreen`
const BITSCREEN_FILENAME = "BITSCREEN_FILENAME"

//   BITSCREEN_PATH     -- path to the bitscreen file, defaults to `.murmuration` in the user home dir
const BITSCREEN_PATH = "BITSCREEN_PATH"

//   BITSCREEN_SOCKET_PORT  -- server socket port of the bitscreen-updater process
const BITSCREEN_SOCKET_PORT = "BITSCREEN_SOCKET_PORT"

// BITSCREEN_LOAD_FROM_FILE -- specify whether to use the bitscreen file for checking cids.
//    Default is to use the bitscreen-updater process (connects to socket port BITSCREEN_SOCKET_PORT)
const BITSCREEN_LOAD_FROM_FILE = "BITSCREEN_LOAD_FROM_FILE"

func IsLoadFromFileEnabled() bool {
	loadFromFile, exists := os.LookupEnv(BITSCREEN_LOAD_FROM_FILE)
	if !exists || loadFromFile == "" {
		loadFromFile = "false"
	}

    return (loadFromFile == "1") || (loadFromFile == "true")
}

func GetBitscreenFilename() string {
	filename, exists := os.LookupEnv(BITSCREEN_FILENAME)
	if !exists || filename == "" {
		filename = "bitscreen"
	}
	return filename
}

func GetSocketPort() string {
	socketPort, exists := os.LookupEnv(BITSCREEN_SOCKET_PORT)
	if !exists || socketPort == "" {
		socketPort = "5555"
	}

    return socketPort
}

// GetPath returns the filepath to the bitscreen file
func GetPath() string {
    fn := GetBitscreenFilename()
	path, exists := os.LookupEnv(BITSCREEN_PATH)
	if !exists || path == "" {
        homeDir, _ := os.UserHomeDir()
        defaultPath := filepath.Join(homeDir, ".murmuration")
		return filepath.Join(defaultPath, fn)
	} else {
	    return filepath.Join(path, fn)
	}
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

// BlockCidFromFile checks for a CID in ./murmuration/bitscreen
func BlockCidFromFile(cidToCheck cid.Cid) bool {
	MaybeCreateBitscreen()
	p := GetPath()

	cidList, err := gabs.ParseJSONFile(p)
	if err != nil {
			panic(err)
	}

	stringList, err := cidList.Children()

	for _, cidString := range stringList {
		string := cidString.Data().(string)

		if string == cidToCheck.String() {
			return true
		}
	}

	return false
}

// BlockCidFromProcess requests the block status of cid from
// the bitscreen-updater process
func BlockCidFromProcess(cidToCheck cid.Cid) bool {
    socketPort := GetSocketPort()
    fmt.Printf("%+v\n", cidToCheck)
    zctx, _ := zmq4.NewContext()
    // Socket to talk to server
    fmt.Printf("Connecting to the server...\n")
    s, _ := zctx.NewSocket(zmq4.REQ)
    s.Connect("tcp://localhost:" + socketPort)

		request := getRequestForCid(cidToCheck)

    fmt.Printf("Sending cid request %s...\n", request)
    s.Send(request, 0)
    responseJSON, _ := s.Recv(0)
		response := getResponseFromJSON(responseJSON)
    fmt.Printf("Received reply [ %s ]\n", responseJSON)
	  log.Printf("dealer received '%d' for cid '%s'", response.reject, cidToCheck.String())

	return response.reject == 1
}

func getRequestForCid(cid cid.Cid) string {
		json := gabs.New()
		json.Set(cid.String(), "cid")

		return json.String()
}

func getResponseFromJSON(responseJSON string) updaterResponse {
		response := updaterResponse{}
		parsed, err := gabs.ParseJSON([]byte(responseJSON))
		if err != nil {
				return response
		}

		if parsed.Exists("error") {
			response.err = parsed.Path("error").Data().(string)

			return response
		}

		response.reject = int(parsed.Path("reject").Data().(float64))
		response.dealCid = parsed.Path("dealCid").Data().(string)
		response.cid = parsed.Path("cid").Data().(string)

		return response
}
