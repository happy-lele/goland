package main

func main() {
	// 回调函数就是通过函数参数将它传到另外一个函数，让它去执行
	DoOperation(1, increase)
	DoOperation02(1, decrease)
}

func DoOperation(y int, f func(int, int) int) {
	f(y, 1)
}

func increase(a, b int) int {
	return a + b
}

func DoOperation02(y int, f func(int, int)) {
	f(y, 1)
}

func decrease(a, b int) {
	println("decrease result is:", a-b)
}
