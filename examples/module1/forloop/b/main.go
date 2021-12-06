package main

import (
	"fmt"
	"sync"
)

func main() {
	// fatal error: all goroutines are asleep - deadlock!
	loopFunc()
}

func loopFunc() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("这样写，会导致循环死锁，因为defer要在loopFunc函数退出时才执行")
		fmt.Println("当第二次循环进来的时候，发现第一次循环的锁还没有释放，就死锁了")
	}
}

func solveLoopFunc()  {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		go func() {
			lock.Lock()
			defer lock.Unlock()
			fmt.Println("defer用来确保，无论后面的代码逻辑是正确还是退出，资源都会被释放")
			fmt.Println("这样写不会导致死锁")
		}()
	}
}


