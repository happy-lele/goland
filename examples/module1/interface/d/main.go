package main

import "fmt"

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func main() {
	var s *Student

	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
	}

	var p People = s

	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
	// s is nil And  p is not nil
	// 当且仅当动态值和动态类型都为 nil 时，接⼝类型值才为 nil
	// p 的动态值是 nil，但是动态类型却是 *Student，是⼀个 nil 指针
}
