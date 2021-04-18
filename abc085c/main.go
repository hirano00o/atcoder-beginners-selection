package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var str string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str = scanner.Text()
	}
	strSlice := strings.Split(str, " ")
	n, _ := strconv.Atoi(strSlice[0])
	y, _ := strconv.Atoi(strSlice[1])
	i, j := 0, 0
	for i = 0; i <= n && i*10000 <= y; i++ {
		for j = 0; j <= n && i*10000+j*5000 <= y; j++ {
			if 10000*i+5000*j+1000*(n-i-j) == y {
				fmt.Printf("%d %d %d\n", i, j, n-i-j)
				return
			}
		}
	}
	fmt.Println("-1 -1 -1")
}
