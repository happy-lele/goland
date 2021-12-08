package main

import "fmt"

var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	// panic: runtime error: invalid memory address or nil pointer dereference
	// nil 指针取值的时候，会panic
	fmt.Println(*p)
}

func main() {
	p, err := foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p)
}