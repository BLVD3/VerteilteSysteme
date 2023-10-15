package main

import (
	"bufio"
	"fmt"
	"os"

	"hhn.de/nvogel1/gochat/server"
)

func main() {
    server, err := server.StartChatServer("127.0.0.1", "3333")
    if err != nil {
        fmt.Println(err)
    }
    bufio.NewReader(os.Stdin).ReadLine()
    server.StopServer()
}
