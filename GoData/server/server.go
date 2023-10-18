package server

import (
	"net"
	"sync"
)

type DataServer struct {
    ip          string
    running     bool
    listener    net.Listener
    data        map[string]string
    lock        sync.RWMutex
}

func NewDataServer(host string, port string) *DataServer {
    return &DataServer{
        ip:         host + ":" + port,
        running:    false,
        data:       make(map[string]string),
    }
}

func (server *DataServer) StartDataServer() error {
    var err error
    server.listener, err = net.Listen("tcp", server.ip)
    if err != nil {
        return err
    }
    go server.acceptConnections()
    return nil
}

func (server *DataServer) acceptConnections() {
    defer server.listener.Close()
    for {

    }
}

func (server *DataServer) handleConnection() {

}
