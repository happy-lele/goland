package main

import (
	"fmt"
)

func main() {
	// map的key不能使用复杂类型，因为没法进行比较
	myMap := make(map[string]string, 10)
	myMap["a"] = "b"
	myFuncMap := map[string]func() int{
		"funcA": func() int { return 1 },
	}
	fmt.Println(myFuncMap)
	f := myFuncMap["funcA"]
	fmt.Println(f())
	value, exists := myMap["a"]
	if exists {
		println(value)
	}
	for k, v := range myMap {
		println(k, v)
	}
}
