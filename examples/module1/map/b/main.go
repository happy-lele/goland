package main

import "fmt"

type Math struct {
	x, y int
}

var m = map[string] Math {
	"foo": Math{2, 3},
}

func main() {
	// 错误原因：对于类似 X = Y 的赋值操作，必须知道 X 的地址，才能够将 Y 的值赋给 X ，
	// 但 go 中的 map 的 value 本身是不可寻址的。
	//m["foo"].x = 4 cannot assign to struct field m["foo"].x in map
	fmt.Println(m["foo"].x)

	math := Math{3, m["foo"].y}
	m["foo"] = math
	fmt.Println(m["foo"])

	// 使⽤临时变量
	tmp := m["foo"]
	tmp.x = 4
	m["foo"] = tmp
	fmt.Println(m["foo"].x)

	// 修改数据结构
	// map[string]*Math
	// m["foo"].x = 4

}
