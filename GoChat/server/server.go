package server

import (
	"encoding/json"
	"net"
	"sync"

	"hhn.de/nvogel1/gochat/shared/messages"
)

type Client struct {
    connection  *net.Conn
}

type ChatServer struct {
    ip          string
    clients     []*Client
    cLock       sync.RWMutex
    ch chan     messages.TextMessage
    running     bool
    listener    net.Listener
}

func StartChatServer(host string, port string) (*ChatServer, error) {
    cs := &ChatServer{
        ip:      host + ":" + port,
        clients: make([]*Client, 0),
        cLock:   sync.RWMutex{},
        ch:      make(chan messages.TextMessage),
        running: true,
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
        if !server.running {
            break
        }
        if err != nil {
            continue
        }
        go server.handleConnection(conn)
    }
}

func (server *ChatServer) handleConnection(conn net.Conn) {
    defer server.StopServer()
    reader := json.NewDecoder(conn)
    client := Client{&conn}
    server.addClient(&client)
    defer server.removeClient(&client)

    var message messages.TextMessage
    for {
        err := reader.Decode(&message)
        if err != nil {
            break
        }
        server.ch <- message
    }
}

func (server *ChatServer) sendMessages() {
    for m := range server.ch {
        b, _ := json.Marshal(m)
        server.cLock.RLock()
        for _, c := range server.clients {
            (*c.connection).Write(b)
        }
        server.cLock.RUnlock()
    }
}

func (server *ChatServer) addClient(client *Client) {
    server.cLock.Lock()
    server.clients = append(server.clients, client)
    server.cLock.Unlock()
}

func (server *ChatServer) removeClient(client *Client) {
    var rIndex int
    for i, c := range server.clients {
        if c == client {
            rIndex = i
            break
        }
    }
    server.cLock.Lock()
    server.clients = append(server.clients[:rIndex], server.clients[rIndex+1:]...)
    server.cLock.Unlock()
}

func (server *ChatServer) StopServer() {
    if !server.running {
        return
    }
    server.running = false
    server.listener.Close()
    close(server.ch)
}
