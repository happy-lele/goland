package main

import (
	"fmt"
	"strings"
)

func main() {
	// 判断字符串中字符是否全都不同 不准使⽤额外的储存结构，且字符串⼩于等于3000
	fmt.Println(isUniqueString("a"))
}

func isUniqueString(s string) bool {
	if strings.Count(s, "") - 1 > 3000 {
		return false
	}
	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}