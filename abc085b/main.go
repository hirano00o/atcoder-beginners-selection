package main

import (
	"fmt"
)

func main() {
	var n, d int
	fmt.Scan(&n)
	m := make(map[int]struct{})
	for i := 0; i < n; i++ {
		fmt.Scan(&d)
		m[d] = struct{}{}
	}
	fmt.Println(len(m))
}
