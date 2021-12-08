package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan int)
	go prod(c)
	go consume(c)
	time.Sleep(time.Second * 10)
}

// 只可以向channel里面发送数据
func prod(ch chan <- int) {
	for  {
		ch <- 1
	}
}

// 只可以从channel里面读取数据
func consume(ch <- chan int) {
	for {
		fmt.Println(<-ch)
	}
}