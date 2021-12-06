package main

import "fmt"

// MyInt1 这是一个新类型
type MyInt1 int

// MyInt2 int的类型别名
type MyInt2 = int

func main() {
	var i int = 0

	// 编译不通过，因为将int类型的变量赋值给MyInt1新类型的变量
	//var i1 MyInt1 = i
	var i1 MyInt1 = MyInt1(i)

	// MyInt2 只是int的别名，类型别名的定义是 =
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}


