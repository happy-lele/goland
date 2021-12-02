package main

import "fmt"

func ReadNoDataFromNoBufCh() {
	noBufCh := make(chan int)
	<- noBufCh
	// fatal error: all goroutines are asleep - deadlock!
	fmt.Println("read from no buffer channel success")
}

func WriteNoBufCh() {
	ch := make(chan int)
	ch <- 1
    // fatal error: all goroutines are asleep - deadlock!
	fmt.Println("write success no block")
}

func ReadNoDataFromBufCh() {
	bufCh := make(chan int, 1)
	<- bufCh
	// fatal error: all goroutines are asleep - deadlock!
	fmt.Println("read from no buffer channel success")
}

func WriteBufChButFull() {
	ch := make(chan int, 1)
	ch <- 100
	ch <- 1
	// fatal error: all goroutines are asleep - deadlock!
	fmt.Println("write success no block")
}

func main()  {
	//ReadNoDataFromNoBufCh()
	//WriteNoBufCh()
	//ReadNoDataFromBufCh()
	WriteBufChButFull()
}
