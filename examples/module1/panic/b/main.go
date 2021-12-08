package  main

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
	// 对于使⽤ := 定义的变量，如果新变量与同名已定义的变量不在同⼀个作⽤域中，
	// 那么 Go 会新定义这个变量 main() 函数⾥的 p 是新定义的变量，会遮住全局变量 p，
	// 导致执⾏到 bar() 时程序，全局变量 p 依然还是 nil，程序随即 Crash。
	p, err := foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p)
}
