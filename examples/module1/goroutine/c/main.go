package main

import (
	"fmt"
	"time"
)

func main() {
	var m = [...]int{1, 2, 3}
	// 各个 goroutine 中输出的 i、v 值都是 for range 循环结束后的 i、v 最终值，
	// ⽽不是各个goroutine启动时的i, v值。可以理解为闭包引⽤，使⽤的是上下⽂环境的值。
    // no
	for i, v := range m {
		go func() {
			 fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second * 2)

	// no
	for i, v := range m {
		go func(int, int) {
			fmt.Println(i, v)
		}(i, v)
	}
	time.Sleep(time.Second * 2)

	// yes
	for i, v := range m {
		go func(i, v int) {
			fmt.Println(i, v)
		}(i, v)
	}
	time.Sleep(time.Second * 2)

	// yes
	for i, v := range m {
		i := i // 这⾥的 := 会重新声明变量，⽽不是重⽤
		v := v
		go func() {
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second * 2)
}
