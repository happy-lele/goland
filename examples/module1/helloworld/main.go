package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// k8s任何程序初始化的时候都会这样去读参数
	name := flag.String("name", "world", "specify the name you want to say hi")
	flag.Parse()
	fmt.Println("os args is:", os.Args)
	fmt.Println("input parameter is:", *name)
	fullString := fmt.Sprintf("Hello %s from Go\n", *name)
	fmt.Println(fullString)
}
