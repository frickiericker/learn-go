package main

import (
	"fmt"
	"sync"
	"time"
)

// Message is a type for the messages passed around in this example.
type Message int

// FinalMessage is sent by producer before exit.
const FinalMessage = Message(3)

// Send represents the sender side of a message channel.
type Send struct {
	MessageChan chan<- Message
}

// Recv represents the receiver side of a message channel.
type Recv struct {
	MessageChan <-chan Message
}

// PairChan creates a message channel and returns its sender and receiver sides.
func PairChan() (Send, Recv) {
	messageChan := make(chan Message)
	return Send{messageChan}, Recv{messageChan}
}

func main() {
	const numReceivers = 3

	prodChan := make(chan Message)
	sendChan := make(chan Send)

	go producer(prodChan)
	go broadcaster(prodChan, sendChan)

	wg := sync.WaitGroup{}
	for i := 0; i < numReceivers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			send, recv := PairChan()
			sendChan <- send
			receiver(recv)
		}()
	}
	wg.Wait()
}

func producer(prodChan chan<- Message) {
	defer close(prodChan)
	for message := Message(0); message <= FinalMessage; message++ {
		prodChan <- message
		time.Sleep(10 * time.Microsecond)
	}
}

func broadcaster(prodChan <-chan Message, sendChan <-chan Send) {
	sends := make([]Send, 0)
	for {
		select {
		case message := <-prodChan:
			for _, send := range sends {
				send.MessageChan <- message
			}
		case send := <-sendChan:
			sends = append(sends, send)
		}
	}
}

func receiver(recv Recv) {
	for message := range recv.MessageChan {
		fmt.Println(&recv, message)
		if message == FinalMessage {
			break
		}
	}
}
