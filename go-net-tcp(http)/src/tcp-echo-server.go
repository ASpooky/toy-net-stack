package main

import (
	"fmt"
	"net"
	"os"

	"go-net-tcp/handler"
)

func main() {
	fmt.Println("process ready.")

	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("[Error]:internal server error")
			os.Exit(1)
		}

		go handler.HandleConnection(conn)
	}
}
