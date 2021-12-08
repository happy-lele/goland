package main

import (
	"fmt"
	"reflect"
)

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func (stu *Student) pSpeak(think string) (talk string) {
	if think == "speak" {
		talk = "pSpeak"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People = Student{}
	think := "speak"
	fmt.Println(reflect.TypeOf(peo), peo.Speak(think))
	// peo.pSpeak undefined (type People has no field or method pSpeak)
	//fmt.Println(peo, peo.pSpeak(think))

	var peo1 People = new(Student)
	fmt.Println(reflect.TypeOf(peo1), peo1.Speak(think))
	// peo1.pSpeak undefined (type People has no field or method pSpeak)
	//fmt.Println(peo1, peo1.pSpeak(think))   奇怪之处

	peo2 := new(Student)
	fmt.Println(reflect.TypeOf(peo2), peo2.Speak(think))
	fmt.Println(reflect.TypeOf(peo2), peo2.pSpeak(think))
}
