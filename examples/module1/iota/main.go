package main

import "fmt"

type Direction int

const (
	North Direction = iota
	East
	South
	West
)
func (d Direction) String() string {
	return [...]string{"North01", "East01", "South01", "West01"}[d]
}

func main() {
	fmt.Println(South)
}
