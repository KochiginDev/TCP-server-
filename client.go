package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	if str == "OK\n" {
		fmt.Println("COOOOL!")
	}
}
