package main

import (
	"fmt"
	"net"
	"strings"
	"os"
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
		rlen, remote, err := conn.ReadFromUDP(message[:])
		if err != nil {
			panic(err)
		}

		data := strings.TrimSpace(string(message[:rlen]))
		if FileExists("/testdata/text.txt") {
			wdata := []byte(data)
			err := os.WriteFile("t.txt",wdata,0644)
			if err != nil {
				panic(err)
			}
		}
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
