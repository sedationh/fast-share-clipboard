package main

import (
	"github.com/google/uuid"
	"net"
	"sync"
)

type PasteConnection struct {
	conn net.Conn
	id   string
}

func NewPasteConnection(conn net.Conn) *PasteConnection {
	return &PasteConnection{
		conn: conn,
		id:   conn.RemoteAddr().String() + "-" + uuid.New().String()[:5],
	}
}

type PasteConnectionList struct {
	connList map[string]*PasteConnection
	lock     sync.Mutex
}

var pasteConnectionList PasteConnectionList = PasteConnectionList{
	connList: make(map[string]*PasteConnection),
}

func (receiver *PasteConnectionList) add(connection *PasteConnection) {
	receiver.lock.Lock()
	receiver.connList[connection.id] = connection
	defer receiver.lock.Unlock()
}
