package main

import "fmt"

func f(n int) (r int) {
	defer func() {
		r += n
		fmt.Println("r = n + 1", r)
		recover()
	}()

	var f func()
	defer f()

	f = func() {
		r += 2
		fmt.Println("r = r + 2", r)
	}
	//defer f()

	fmt.Println("r = finally", n)
	return n + 1
}

func main() {
	fmt.Println(f(3))
}


