package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan int, 10)
	// producer
	for i := 0; i < 10; i++ {
		messages <- i
	}

	go func() {
		for _ = range messages {
			fmt.Println("aaaaaa")
			select {
			case i := <-messages:
				fmt.Println(i)
				break
			}
		}
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("main process exit!")
}
