package main

import (
	"context"
	"golang.design/x/clipboard"
	"log"
	"net"
)

func ClientRunner() {
	conn, err := net.Dial("tcp", config.Host+":"+config.Port)
	if err != nil {
		return
	}
	pasteConnection := NewPasteConnection(conn)
	defer conn.Close()

	err = clipboard.Init()
	if err != nil {
		return
	}

	clipboardWatch := clipboard.Watch(context.Background(), clipboard.FmtText)
	lastContent := clipboard.Read(clipboard.FmtText)

	connReadCh := make(chan []byte, 1)

	go func() {
		for {
			content, err := pasteConnection.Read()
			if err != nil {
				return
			}
			connReadCh <- content
		}
	}()

	for {
		select {
		case content := <-clipboardWatch:
			log.Printf("client clipboardWatch write: %s\nlastContent: %s\n", string(content), string(lastContent))
			if string(lastContent) == string(content) {
				continue
			}
			lastContent = content
			pasteConnection.Send(content)
		case content := <-connReadCh:
			log.Printf("client received: %s\nlastContent: %s\n", string(content), string(lastContent))
			if string(lastContent) == string(content) {
				continue
			}
			clipboard.Write(clipboard.FmtText, content)
			lastContent = content
		}
	}
}
