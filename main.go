package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	ln, _ := net.Listen("tcp", ":8080")

	for {
		conn, _ := ln.Accept()
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 4096)
	// conn.Read(buffer)
	// fmt.Println(string(buffer))
	fmt.Println("Request")

	file, _ := os.Open("index.html")
	defer file.Close()
	file.Read(buffer)

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	conn.Write(buffer)
	conn.Close()
}
