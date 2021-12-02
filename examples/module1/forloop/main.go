package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fullString := "hello world"
	fmt.Println(fullString)
	// 需要string(c)才能取到字符
	for i, c := range fullString {
		fmt.Println(i, string(c))
	}

}
