package main

import (
	"encoding/json"
	"fmt"
	"net"

	"hhn.de/nvogel1/gochat/shared/messages"
)

const (
	CONN_HOST = "127.0.0.1"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func startServer(connType string, host string, port string) error {
	ln, err := net.Listen(connType, host + ":" + port)
	if err != nil {
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Connection Failed: ", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	reader := json.NewDecoder(conn)
	for {
		var m map[string] any
		decErr := reader.Decode(&m)
		if decErr != nil {
			fmt.Println("Decoding Error: ", decErr)
			break
		}
		//Handle the map
		t := messages.GetMessageType(m)
		fmt.Println(t)
	}
	conn.Close()
}

func main() {
	startServer(CONN_TYPE, CONN_HOST, CONN_PORT)
}

