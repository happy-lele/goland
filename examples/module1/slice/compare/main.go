package main

func main() {
	//invalid operation: [1]int{...} == [2]int{...} (mismatched types [1]int and [2]int)
	//fmt.Println([...]int{1} == [2]int{1})

	// invalid operation: []int{...} == []int{...} (slice can only be compared to nil)
	//fmt.Println([]int{1} == []int{1})

}
