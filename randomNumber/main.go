package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	str := ""
	for i := 0; i < 4; i++ {
		str += fmt.Sprint(randInt())
	}
	fmt.Println(str)
}

func randInt() int {
	return rand.Intn(9)
}
