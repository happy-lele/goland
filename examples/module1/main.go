package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name" protobuf:"bytes"`
}

func main() {
	my := Person{Name: "zhang le"}

	myType := reflect.TypeOf(my)
	name := myType.Field(0)

	tag := name.Tag.Get("protobuf")
	fmt.Println(myType)
	fmt.Println(name)
	fmt.Println(tag)
}

