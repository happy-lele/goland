package main

import "fmt"


/*func noSwap(a int ,b int) {
	var temp int
	temp = a
	a = b
	b = temp
}
*/

func canSwap(pa *int, pb *int) {
	var temp int
	// 直接操作指针
	temp = *pa // temp = main::a
	*pa = *pb  // main::a = main::b
	*pb = temp // main::b = temp
}


func main() {
	var a int = 10
	var b int = 20

	canSwap(&a, &b)
	fmt.Println("a = ", a, " b = ", b)

	var p *int
	p = &a

	fmt.Println(&a)
	fmt.Println(p)

	var pp **int //二级指针

	pp = &p

	fmt.Println(&p)
	fmt.Println(pp)
}
