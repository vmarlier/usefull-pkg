package main

//// Use goroutine with counter
import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//// Counted concurreny goroutine
	var wg sync.WaitGroup // Init wait group (counter)
	wg.Add(1)             // Increment wait group

	//// Creation function to call count function to decrement the counter (wait group)
	//// The decrement is performed hereunder because is not count function responsability to ensure the decrement
	go func() {
		count("user")
		wg.Done() // Decrement the counter
	}() // Execute directly the function

	wg.Wait() // Block the function until the counter is 0
}

func count(something string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, something)
		time.Sleep(time.Millisecond * 200)
	}
}
