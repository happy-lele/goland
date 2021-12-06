package main

import "fmt"

type Person struct {
	age int
}

func main() {
	person := &Person{28}
	// 28 会把 28 缓存在栈中，等到最后执⾏该defer语句的时候取出
	defer fmt.Println(person.age)
	// 28 这个地址指向的结构体没有被改变
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)
	// 29
	defer func() {
		fmt.Println(person.age)
	}()

	person = &Person{29}
}