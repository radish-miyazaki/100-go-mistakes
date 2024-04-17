package main

import (
	"fmt"
	"runtime"
)

func consumeMessages() {
	for {
		msg := receiveMessage()
		storeMessageType(msg)
	}
}

func getMessageType(msg []byte) []byte {
	return msg[:5]
}

func getMessageTypeWithCopy(msg []byte) []byte {
	msgType := make([]byte, 5)
	copy(msgType, msg)
	return msgType
}

func receiveMessage() []byte {
	return make([]byte, 1_000_000)
}

func storeMessageType(msg []byte) {}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
