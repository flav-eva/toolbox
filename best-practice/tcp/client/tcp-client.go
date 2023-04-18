package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		_, err = conn.Write([]byte("hello"))
		if err != nil {
			panic(err)
		}

		time.Sleep(3 * time.Second)

		buf := make([]byte, 1024)
		// _, err = conn.Read(buf)
		// if err != nil {
		// 	panic(err)
		// }

		fmt.Println(string(buf))
	}

}
