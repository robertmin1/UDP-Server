package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 3000,
		IP:   net.ParseIP("0.0.0.0"),
	})
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	fmt.Printf("server listening %s\n", conn.LocalAddr().String())
	
	for {
		message := make([]byte, 20)

		if FileExists("testdata/text.txt") {
			_, err := exec.Command("/bin/bash", "send-to-sever.bash").Output()
			if err != nil {
				panic(err)
			}
			rlen, _, err := conn.ReadFromUDP(message[:])
			data := strings.TrimSpace(string(message[:rlen]))
			wdata := []byte(data)
			err1 := os.WriteFile("testdata/text.txt",wdata,0644)
			if err1 != nil {
				panic(err1)
			}
			os.Exit(0)
		}

		rlen, remote, err := conn.ReadFromUDP(message[:])
		if err != nil {
			panic(err)
		}

		data := strings.TrimSpace(string(message[:rlen]))
		
		fmt.Printf("received: %s from %s\n", data, remote)
	}
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
