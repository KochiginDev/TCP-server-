package main

import (
	"fmt"
	"net"
)

func Connection(conn net.Conn) {
	defer conn.Close()
	message := "OK\n"
	_, err := conn.Write([]byte(message))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	fd, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fd.Close()
	for {
		con, err := fd.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go Connection(con)
	}
}
