package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer fun is called")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("a panic is triggered")

}
