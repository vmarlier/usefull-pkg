///////// Final test //////////
package main

import (
	"fmt"
	"time"
)

func main() {
	// Create and start a broker:
	w := NewWorker()
	go w.StartConversation()
	// Create and subscribe 3 clients:
	clientFunc := func(id int) {
		msgCh := w.Subscribe()
		for {
			fmt.Printf("Client %d got message: %v\n", id, <-msgCh)
		}
	}
	for i := 0; i < 3; i++ {
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

type Worker struct {
	pubChan chan interface{}
	subChan chan chan interface{}
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

//// Simple way to use goroutine to unblock runtime
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	go count("user")
//	count("message")
//}
//
//func count(something string) {
//	for i := 1; i <= 5; i++ {
//		fmt.Println(i, something)
//		time.Sleep(time.Millisecond * 200)
//	}
//}
//// Use goroutine with counter
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//func main() {
//	//// Counted concurreny goroutine
//	var wg sync.WaitGroup // Init wait group (counter)
//	wg.Add(1)             // Increment wait group
//
//	//// Creation function to call count function to decrement the counter (wait group)
//	//// The decrement is performed hereunder because is not count function responsability to ensure the decrement
//	go func() {
//		count("user")
//		wg.Done() // Decrement the counter
//	}() // Execute directly the function
//
//	wg.Wait() // Block the function until the counter is 0
//}
//
//func count(something string) {
//	for i := 1; i <= 5; i++ {
//		fmt.Println(i, something)
//		time.Sleep(time.Millisecond * 200)
//	}
//}
//// Use goroutine with channel to communicate between several goroutine
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	cn := make(chan string) // User make function to create channel
//	go count("message", cn)
//
//	//// Get all messages (explanation version)
//	//for {
//	//	msg, open := <-cn // Get the message to the channel
//
//	//	// Stop receiveing messages if channel is closed
//	//	if !open {
//	//		break
//	//	}
//
//	//	fmt.Println(msg)
//	//}
//
//	// Better way to close receiver
//	for msg := range cn {
//		fmt.Println(msg)
//	}
//}
//
//func count(something string, cn chan string) {
//	for i := 1; i <= 5; i++ {
//		cn <- something // Send message over the channel
//		time.Sleep(time.Millisecond * 200)
//	}
//
//	close(cn) // Close channel to avoid "deadlock" in async runtime
//}
//// Channel explanations
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	cn1 := make(chan string) // User make function to create channel 1
//	cn2 := make(chan string) // User make function to create channel 2
//
//	// Init Sender message function
//	go func() {
//		for {
//			cn1 <- "This message is sent every 500ms" // Sending message
//			time.Sleep(time.Millisecond * 500)
//		}
//	}()
//
//	go func() {
//		for {
//			cn2 <- "This message is sent every 2s" // Sending message
//			time.Sleep(time.Second * 2)
//		}
//	}()
//
//	// Init Receiver message in main
//	for {
//		// Select allow to output message once is received
//		select {
//		case msg1 := <-cn1:
//			fmt.Println(msg1)
//		case msg2 := <-cn2:
//			fmt.Println(msg2)
//		}
//	}
//}
///////// Associate all concepts from the begining
// goroutine using only one channel
import "fmt"

func main() {
	c := make(chan int)
	for i := 1; i <= 5; i++ {
		go func(i int) {
			for v := range c {
				fmt.Printf("count %d from goroutine #%d\n", v, i)
			}
		}(i)
	}
	for i := 1; i <= 25; i++ {
		c <- i
	}
	close(c)
}

//// goroutine with multi channels. Simulate a queue. (IN PROGRESS)
//import (
//	"fmt"
//)
//
//func main() {
//	numOfUsers := make(chan int, 5) // User make function to create channel 1
//	n := make() // User make function to create channel 2
//
//    // Create worker to dispatch messages in proper users channel
//
//
//	// Create function to count the number of user in the conversationId
//    func list_user(n int) int {
//        if n < 1 {
//            return "There is no user in the conversation"
//        }
//
//        return "5"
//    }
//
//}
