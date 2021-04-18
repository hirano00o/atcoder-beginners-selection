package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	words := [4]string{"remaerd", "maerd", "resare", "esare"}
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str = scanner.Text()
	}
	ss := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
	str = string(ss)
	for i := 0; i < 4; i++ {
		s := strings.TrimPrefix(str, words[i])
		if s != str {
			str = s
			i = 0
		}
	}
	if len(str) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
