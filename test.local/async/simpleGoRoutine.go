package main

// Simple way to use goroutine to unblock runtime
import (
	"fmt"
	"time"
)

func main() {
	go count("user")
	count("message")
}

func count(something string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, something)
		time.Sleep(time.Millisecond * 200)
	}
}
