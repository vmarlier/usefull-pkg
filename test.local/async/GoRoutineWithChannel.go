package main

//// Use goroutine with channel to communicate between several goroutine
import (
	"fmt"
	"time"
)

func main() {
	cn := make(chan string) // User make function to create channel
	go count("message", cn)

	//// Get all messages (explanation version)
	//for {
	//	msg, open := <-cn // Get the message to the channel

	//	// Stop receiveing messages if channel is closed
	//	if !open {
	//		break
	//	}

	//	fmt.Println(msg)
	//}

	// Better way to close receiver
	for msg := range cn {
		fmt.Println(msg)
	}
}

func count(something string, cn chan string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		cn <- something // Send message over the channel
		time.Sleep(time.Millisecond * 200)
	}

	close(cn) // Close channel to avoid "deadlock" in async runtime
}
