package main

import (
	"context"
	"fmt"
	"golang.design/x/clipboard"
	"time"
)

func TestMain() {
	err := clipboard.Init()
	if err != nil {
		return
	}

	clipboardCh := clipboard.Watch(context.Background(), clipboard.FmtText)

	for i := 0; i < 5; i++ {
		i := i
		go func() {
			time.Sleep(time.Duration(2*i) * time.Second)
			clipboard.Write(clipboard.FmtText, []byte("你好，world 哈哈哈"+fmt.Sprint(i)))
		}()
	}

	for {
		clipboardText := string(<-clipboardCh)
		fmt.Println(clipboardText)
	}
}
