package processenum


import (
	"github.com/shirou/gopsutil/v3/process"
	"os"
	"fmt"
)

type Process struct {
	ProcessName	string
	ProcessFile	[]string
	ProcessUser	string
	ProcessIdentifier	int
}

func KillProcess(pid int) {
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err)
	}
	err = process.Kill()
	if err != nil {
		fmt.Println("[!] Could not kill process: " + err.Error())
	} else {
		fmt.Println("[+] Process killed!")
	}
}

func EnumerateProcesses() []Process {
	procList, err := process.Processes()
	if err != nil {
		fmt.Println(err)
	}
	processes := make([]Process, len(procList))
	for i, process := range procList {
		exe, err := process.CmdlineSlice()
		if err != nil {
			fmt.Println(err)
		}
		name, err := process.Name()
		if err != nil {
			fmt.Println(err)
		}
		user, err := process.Username()
		if err != nil {
			fmt.Println(err)
		}
		processes[i] = Process{
			ProcessName:	name,
			ProcessFile:	exe,		
			ProcessUser:	user,
			ProcessIdentifier:	int(process.Pid),
		}
	}
	return processes
}
