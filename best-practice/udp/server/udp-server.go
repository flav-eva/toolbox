package main

import (
	"fmt"
	"net"
)

func main() {
	raddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("listening on", raddr)
	conn, err := net.ListenUDP("udp", raddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// receive message
	buf := make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println("n:", n)

	// send message
	n, err = conn.WriteToUDP([]byte("hello"), addr)
	if err != nil {
		panic(err)
	}
}
