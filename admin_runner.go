package main

import (
	"fmt"
	"log"
	"net"
)

func pasteConnectionMsgHandler(pasteConnection *PasteConnection) {
	defer pasteConnection.conn.Close()
	for {
		content, err := pasteConnection.Read()
		if err != nil {
			return
		}
		if err != nil {
			return
		}
		log.Println("admin received", string(content))

		// 向其他 pasteConnection 发送消息
		for _, pasteConnection := range pasteConnectionList.connList {

			log.Printf("admin send to %s", pasteConnection.id)

			pasteConnection.Send(content)
		}
	}
}

func pasteConnectionHandler(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		pasteConnection := NewPasteConnection(conn)
		pasteConnectionList.add(pasteConnection)

		// 处理 message
		go pasteConnectionMsgHandler(pasteConnection)
	}
}

func AdminRunner() {
	if config.Role != "admin" {
		return
	}

	log.Printf("admin local host: %v", GetOutboundIP())

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", config.Port))
	if err != nil {
		panic(err)
	}

	// 处理关闭 listener
	go pasteConnectionHandler(listener)
}

// Get preferred outbound ip of this machine
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
