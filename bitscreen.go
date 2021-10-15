package bitscreen

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"github.com/ipfs/go-cid"
	"github.com/pebbe/zmq4"

)

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
	f, err := os.OpenFile(p, os.O_RDONLY, os.ModePerm)
	if err != nil {
		sigterm(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	sigterm(s.Err())

	for {
		b := s.Scan()
		if b {
			if s.Text() == cidToCheck.String() {
				fmt.Printf("Deals for CID %s are not welcome.\r\n", cidToCheck.String())
				return true
			}
		} else {
			break
		}
	}
	return false
}

// BlockCidFromProcess requests the block status of cid from
// the bitscreen-updater process
func BlockCidFromProcess(cidToCheck cid.Cid) bool {
    socketPort := GetSocketPort()

    zctx, _ := zmq4.NewContext()

    // Socket to talk to server
    fmt.Printf("Connecting to the server...\n")
    s, _ := zctx.NewSocket(zmq4.REQ)
    s.Connect("tcp://localhost:" + socketPort)

    fmt.Printf("Sending cid request %s...\n", cidToCheck.String())
    s.Send(cidToCheck.String(), 0)
    cidBlocked, _ := s.Recv(0)
    fmt.Printf("Received reply [ %s ]\n", cidBlocked)
	log.Printf("dealer received '%s' for cid '%s'", cidBlocked, cidToCheck.String())

	return cidBlocked == "1"
}
