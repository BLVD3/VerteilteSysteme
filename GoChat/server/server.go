package server

import (
	"encoding/json"
	"errors"
	"net"
	"sync"

	"hhn.de/nvogel1/gochat/shared/messages"
)

type Client struct {
    conn *net.Conn
    name string
}

type ChatServer struct {
    ip string
    clients []*Client
    cLock sync.RWMutex
    ch chan messages.ReceiveTextMessage
    close bool
    listener net.Listener
}

func StartChatServer(host string, port string) (*ChatServer, error) {
    cs := &ChatServer{
        ip:      host + ":" + port,
        clients: make([]*Client, 0),
        cLock:   sync.RWMutex{},
        ch:      make(chan messages.ReceiveTextMessage),
        close:   false,
    }
    ln, err := net.Listen("tcp", cs.ip)
    if err != nil {
        return nil, err
    }
    cs.listener = ln
    go cs.acceptConnections()
    go cs.sendMessages()
    return cs, nil;
}

func (server *ChatServer) acceptConnections() {
    defer server.listener.Close()
    defer close(server.ch)
    for {
        conn, err := server.listener.Accept()
        if server.close {
            break
        }
        if err != nil {
            continue
        }
        go server.handleConnection(conn)
    }
}

func (server *ChatServer) handleConnection(conn net.Conn) {
    defer conn.Close()
    reader := json.NewDecoder(conn)
    server.addClient(&conn, reader)

}

func (server *ChatServer) addClient(conn *net.Conn, reader *json.Decoder) error {
    (*conn).Write([]byte(messages.NameRequestMessageString))

    var m map[string]any
    err := reader.Decode(&m)
    if err != nil {
        return err
    }
    var t = messages.GetMessageType(&m)
    if t != messages.NameResponse {
        return errors.New("failed to read name from client: Expected NameResponseMessage got " + t.String())
    }
    message, err := messages.GetNameResponseMessage(&m)
    if err != nil {
        return err;
    }
    client := Client{conn, message.Name}
    server.cLock.Lock()
    server.clients = append(server.clients, &client)
    server.cLock.Unlock()
    return nil
}

func (server *ChatServer) sendMessages() {
    for m := range server.ch {
        b, _ := json.Marshal(m)
        server.cLock.RLock()
        for _, c := range server.clients {
            (*c.conn).Write(b)
        }
        server.cLock.RUnlock()
    }
}
