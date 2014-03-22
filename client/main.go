package main

import (
	"fmt"
	"net/http"
	"os"
    "encoding/json"
    "strconv"
    "strings"
)

type JSONPayload map[string]interface{}

func (r JSONPayload) String() (s string) {
    b, err := json.Marshal(r)
    if err != nil {
        s = ""
        fmt.Printf("Error marshalling JSON: %s\n", err)
        return
    }
    s = string(b)
    return
}

func usage() {
	fmt.Printf("Grooveshark Extension command line controller\n")
	fmt.Printf("Support commands:\n")
	fmt.Printf("\tstart\n")
	fmt.Printf("\tstop\n")
	fmt.Printf("\ttoggle\n")
	fmt.Printf("\tnext\n")
	fmt.Printf("\tprev\n")
	fmt.Printf("\tvolume #\n")
}

func sendCommand(jsonPayload JSONPayload) error {
    _, err := http.Post("http://localhost:4555", "application/json", strings.NewReader(jsonPayload.String()))
    return err
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	if os.Args[1] == "start" {
        sendCommand(JSONPayload{"command": "start"})
	} else if os.Args[1] == "stop" {
        sendCommand(JSONPayload{"command": "stop"})
	} else if os.Args[1] == "toggle" {
        sendCommand(JSONPayload{"command": "toggle"})
	} else if os.Args[1] == "next" {
        sendCommand(JSONPayload{"command": "next"})
	} else if os.Args[1] == "prev" {
        sendCommand(JSONPayload{"command": "prev"})
	} else if os.Args[1] == "volume" {
        if len(os.Args) < 3 {
            fmt.Printf("Number not supplied to volume command\n")
            return
        }
        volume, err := strconv.Atoi(os.Args[2])
        if err != nil {
            fmt.Printf("Error converting volume to number: %s\n", err)
            return
        }
        if volume < 0 || volume > 100 {
            fmt.Printf("Volume < 0 || > 100: %d: %s\n", volume, err)
        }
        sendCommand(JSONPayload{"command": "volume", "volume": volume})
	} else {
		fmt.Printf("No command provided\n\n")
		usage()
	}
}
