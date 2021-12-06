package main

import (
	"fmt"
)

func main() {
	str := "a string oop"
	pointer := &str
	anotherString := *&str
	fmt.Println(str)
	fmt.Println(pointer)
	fmt.Println(anotherString)

	str = "changed string"
	fmt.Println(str)
	fmt.Println(pointer)
	fmt.Println(anotherString)

	para := ParameterStruct{Name: "aaa"}
	fmt.Println(para)

	changeParameter(&para, "bbb")
	fmt.Println(para)

	cannotChangeParameter(para, "ccc")
	fmt.Println(para)

	cannotChangeParameter02(&para, "ddd")
	fmt.Println(para)
}

type ParameterStruct struct {
	Name string
}

func changeParameter(para *ParameterStruct, value string) {
	// 操作的是指针
	para.Name = value
}

func cannotChangeParameter(para ParameterStruct, value string) {
	para.Name = value
}

func cannotChangeParameter02(para *ParameterStruct, value string) {
	a := *para
	// a只是struct的一个拷贝
	a.Name = value
}
