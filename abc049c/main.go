package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	words := [4]string{"dream", "dreamer", "erase", "eraser"}
	fmt.Scanf("%s", &str)
	for {
		flag := false
		for i := range words {
			if strings.HasSuffix(str, words[i]) {
				str = str[:len(str)-len(words[i])]
				flag = true
				break
			}
		}
		for i := range words {
			if str == words[i] {
				fmt.Println("YES")
				return
			}
		}
		if !flag {
			fmt.Println("NO")
			return
		}
	}
}
