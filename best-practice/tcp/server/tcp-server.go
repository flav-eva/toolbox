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
	fmt.Println("listening on", tcpAddr)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	// listen incoming connection
	conn, err := listener.Accept() // 这里会阻塞，等待客户端连接
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		// send message
		_, err = conn.Write([]byte("world!"))
		if err != nil {
			panic(err)
		}

		time.Sleep(5 * time.Second)

		// receive message
		buf := make([]byte, 1024)
		_, err = conn.Read(buf) // 这里读取客户端发送来的请求，如果客户端关闭，这里会抛错
		if err != nil {
			panic(err)
		}
		fmt.Println(string(buf))
	}
}
