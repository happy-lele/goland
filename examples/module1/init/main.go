package main

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Println("last run. . . ")

	a := errors.New("not found")

	fmt.Println(a)

}
