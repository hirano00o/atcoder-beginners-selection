package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func convertInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func abs(x, y int) int {
	d := x - y
	if d < 0 {
		return -d
	}
	return d
}

func main() {
	var N int
	if scanner.Scan() {
		N = convertInt(scanner.Text())
	}

	prevT, prevX, prevY := 0, 0, 0
	for i := 0; i < N; i++ {
		var t, x, y int
		if scanner.Scan() {
			line := scanner.Text()
			s := strings.Split(line, " ")
			tmp := make([]int, 3)
			for j := range s {
				tmp[j] = convertInt(s[j])
			}
			t = tmp[0]
			x = tmp[1]
			y = tmp[2]
		}
		dt := t - prevT
		dist := abs(x, prevX) + abs(y, prevY)
		if dt < dist || dist%2 != dt%2 {
			fmt.Println("No")
			return
		}
		prevT, prevX, prevY = t, x, y
	}
	fmt.Println("Yes")
}
