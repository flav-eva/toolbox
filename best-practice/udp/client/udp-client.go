package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	raddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("dial udp on", raddr)
	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("dial udp on", raddr)

	for {
		// send message
		n, err := conn.Write([]byte("hello !"))
		if err != nil {
			panic(err)
		}

		// receive message
		fmt.Println("string(buf[:n])")
		buf := make([]byte, 1024)
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			panic(err)
		}
		fmt.Println("string(buf[:n])")
		fmt.Println(string(buf[:n]))
		fmt.Println("arrived from", addr)

		time.Sleep(5 * time.Second)
	}
}
