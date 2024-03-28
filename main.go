package main

import (
	"log"
	ps "github.com/mitchellh/go-ps"
	"os/exec"
	"time"
	"strings"
)

func killQBitTorrent() {
    app := "taskkill.exe"
    arg0 := "/f"
    arg1 := "/im"
    arg2 := "qbittorrent.exe"

    cmd := exec.Command(app, arg0, arg1, arg2)
    stdout, err := cmd.Output()

    if err != nil {
        log.Println(err.Error())
        return
    }

    // Print the output
    log.Println(string(stdout))
}

func main() {
	for {
		processList, err := ps.Processes()
	    if err != nil {
	        log.Println("ps.Processes() Failed, are you using windows?")
	        return
	    }

	    foundVpn := false
	    foundTorrent := false

	    for x := range processList {
	        var process ps.Process
	        process = processList[x]
	        log.Printf("%d\t%s\n", process.Pid(), process.Executable())

	        if strings.EqualFold(process.Executable(), "wireguard.exe") {
	        	foundVpn = true
	        }
	        if strings.EqualFold(process.Executable(), "qbittorrent.exe") {
	        	foundTorrent = true
	        }
	    }

	    if foundVpn && foundTorrent {
	    	killQBitTorrent()
	    	log.Println("QBitTorrent stopped")
	    } else {
	    	log.Println("Wireguard or QBitTorrent not running")
	    }

		time.Sleep(10 * time.Second) 
	}}