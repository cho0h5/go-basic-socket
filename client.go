package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", ":8080")

	go func() {
		data := make([]byte, 1)

		for {
			conn.Read(data)
			fmt.Println(data)
		}
	}()

	for {
		var s string
		fmt.Scanln(&s)
		conn.Write([]byte(s))
	}
}
