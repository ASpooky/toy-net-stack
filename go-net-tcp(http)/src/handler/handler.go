package handler

import (
	"fmt"
	"io"
	"net"
)

func HandleConnection(conn net.Conn) {

	defer conn.Close()

	fmt.Println("=====get request=====")

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)

	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println("読み取りエラー:", err)
		return
	}

	fmt.Println(string(buf[:n]))

	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/html; charset=utf-8\r\n" +
		"\r\n" +
		"<h1>Hello, toy-net-stack!</h1><p>自作TCPサーバーからの返信です。</p>"

	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("書き込みエラー", err)
		return
	}

	fmt.Println("====return response====")
}
