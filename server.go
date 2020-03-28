package main

import (
	"log"
	"net"
	"time"
)

func main() {
	ln, _ := net.Listen("tcp", ":8080")
	log.Println("Started Server")

	for {
		conn, _ := ln.Accept()
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	log.Println("Connected")
	buffer := make([]byte, 1)
	buffer[0] = 0

	for {
		_, err := conn.Write(buffer)
		if err != nil {
			log.Println("Closed")
			break
		}

		buffer[0]++
		time.Sleep(time.Second)
	}
	conn.Close()
}
