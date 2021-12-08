package main

import "fmt"

func main() {
	// 求翻转后的字符串
	fmt.Println(reverseString("abcdef"))
}

func reverseString(s string) (string, bool)  {
	str := []rune(s)
	l := len(s)

	if l > 5000 {
		return s, false
	}

	for i := 0; i < l / 2; i++ {
		str[i], str[l-1-i] = str[l-1-i], str[i]
	}
	return string(str), true
}
