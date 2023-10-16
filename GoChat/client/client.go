package client

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"hhn.de/nvogel1/gochat/shared/messages"
)

type ChatInstance struct {
    ip string
    conn net.Conn
    running bool
    userName string
    sendChannel chan messages.TextMessage
    ReceiveChannel chan messages.TextMessage
}

func ConnectToChatServer(host string, port string, name string) (*ChatInstance, error) {
    conn, err := net.Dial("tcp", host + ":" + port)
    if err != nil {
        return nil, err
    }
    var chat = ChatInstance {
        ip: host + ":" + port,
        conn: conn,
        running: true,
        userName: name,
        sendChannel: make(chan messages.TextMessage),
        ReceiveChannel: make(chan messages.TextMessage),
    }
    go chat.listenForMessages()
    go chat.sendMessages()
    return &chat, nil
}

func (instance *ChatInstance) listenForMessages() {
    reader := json.NewDecoder(instance.conn)
    var message messages.TextMessage
    var err error
    for {
        err = reader.Decode(&message)
        if err != nil {
            break
        }
        instance.ReceiveChannel <- message
    }
}

func (instance *ChatInstance) sendMessages() {
    for m := range instance.sendChannel {
        raw, err := json.Marshal(m)
        if err != nil {
            fmt.Println("Could not Marshal message: " + m.String())
            continue
        }
        instance.conn.Write(raw)
    }
}

func (instance *ChatInstance) Disconnect() {
    instance.running = false
    instance.conn.Close()
    close(instance.ReceiveChannel)
    close(instance.sendChannel)
}

func (instace *ChatInstance) SendMessage(text string) bool {
    if !instace.running {
        return false
    }
    var message messages.TextMessage = messages.TextMessage{
        Sender: instace.userName,
        Time: time.Now().UnixMilli(),
        Text: text,
    }
    instace.sendChannel <- message
    return true
}
