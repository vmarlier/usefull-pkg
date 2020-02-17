package main

import (
	"fmt"
	"time"
)

type Worker struct {
	pubChan chan interface{}
	subChan chan chan interface{}
}

func main() {
	// Create and start a broker:
	w := NewWorker()
	go w.StartConversation()

	// Create and subscribe 3 clients:
	clientFunc := func(id int) {
		msgCh := w.Subscribe()
		for {
			fmt.Printf("Client %d got message: %v\n", id, <-msgCh)
			fmt.Println("ok")
		}
	}
	for i := 0; i < 50; i++ {
		go clientFunc(i)
	}

	// Start publishing messages:
	go func() {
		for msgId := 0; msgId < 1; msgId++ {
			w.Publish(fmt.Sprintf("msg#%d", msgId))
			time.Sleep(300 * time.Millisecond)
		}
	}()

	time.Sleep(time.Second)
}

func NewWorker() *Worker {
	return &Worker{
		pubChan: make(chan interface{}, 1),
		subChan: make(chan chan interface{}, 1),
	}
}

func (wk *Worker) StartConversation() {
	subs := map[chan interface{}]struct{}{}
	for {
		select {
		case msgCh := <-wk.subChan:
			subs[msgCh] = struct{}{}
		case msg := <-wk.pubChan:
			for msgCh := range subs {
				// msgCh is buffered, use non-blocking send to protect the broker:
				select {
				case msgCh <- msg:
				default:
				}
			}
		}
	}
}

func (wk *Worker) Subscribe() chan interface{} {
	msgCh := make(chan interface{}, 5)
	wk.subChan <- msgCh
	return msgCh
}

func (wk *Worker) Publish(msg interface{}) {
	wk.pubChan <- msg
}
