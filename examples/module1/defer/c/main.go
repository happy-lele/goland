package main

import "fmt"

type Person struct {
	age int
}

func main() {
	person := &Person{28}
	// 28
	defer fmt.Println(person.age)
	// 29 缓存的是person的地址
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)
	// 29 闭包引用
	defer func() {
		fmt.Println(person.age)
	}()

	person.age = 29
}
