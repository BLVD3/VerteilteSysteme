package main

import (
	"bufio"
	"fmt"
	"os"

	"hhn.de/nvogel1/gochat/client"
	"hhn.de/nvogel1/gochat/shared/messages"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your name: ")
    name, _ := reader.ReadString('\n')
	server, err := client.ConnectToChatServer("127.0.0.1", "3333", name[:len(name)-1])
	if err != nil {
		fmt.Println(err)
		return
	}
    go printMessageRoutine(&server.ReceiveChannel)
	defer server.Disconnect()
	for {
        input, _ := reader.ReadString('\n')
        if !server.SendMessage(input[:len(input)-1]) {
            break
        }
	}
}

func printMessageRoutine(ch *chan messages.TextMessage) {
    for m := range *ch {
        fmt.Println(m.String())
    }
}

