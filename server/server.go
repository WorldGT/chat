package main

import (
	"chat_socket/server/internal"
	"fmt"
	"net"
)

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:8085")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	defer listen.Close()
	hub := internal.NewHub()
	go hub.Run()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err", err)
		}
		go internal.ServerWs(conn, hub)
	}

}
