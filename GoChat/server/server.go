package server

import (
	"encoding/json"
	"net"
	"sync"

	"hhn.de/nvogel1/gochat/shared/messages"
)

type Client struct {
    conn net.Conn
    name string
}

type ChatServer struct {
    ip string
    clients []*Client
    cLock sync.RWMutex
    ch chan messages.SendTextMessage
    close bool
    listener net.Listener
}

func StartChatServer(host string, port string) (*ChatServer, error) {
    cs := &ChatServer{
        ip:      host + ":" + port,
        clients: make([]*Client, 0),
        cLock:   sync.RWMutex{},
        ch:      make(chan messages.SendTextMessage),
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
    server.addClient(conn)

}

func (server *ChatServer) addClient(conn net.Conn) error {

}

func (server *ChatServer) sendMessages() {
    for m := range server.ch {
        b, _ := json.Marshal(m) //TODO Improve Marshal
        server.cLock.RLock()
        for _, c := range server.clients {
            c.conn.Write(b)
        }
        server.cLock.RUnlock()
    }
}



