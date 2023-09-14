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

func (receiver *PasteConnection) Send(content []byte) {
	contentLen := len(content)
	_, err := receiver.conn.Write(Int64ToBytes(int64(contentLen)))
	if err != nil {
		return
	}
	_, err = receiver.conn.Write(content)
	if err != nil {
		return
	}
}

func (receiver *PasteConnection) Read() ([]byte, error) {
	contentLenInfo := make([]byte, headerLen)
	_, err := receiver.conn.Read(contentLenInfo)
	if err != nil {
		return nil, err
	}
	contentLen := BytesToInt64(contentLenInfo)
	content := make([]byte, contentLen)
	_, err = receiver.conn.Read(content)
	if err != nil {
		return nil, nil
	}
	return content, nil
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
