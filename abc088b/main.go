package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var aStr string
	if scanner.Scan() {
		_ = scanner.Text()
	}
	if scanner.Scan() {
		aStr = scanner.Text()
	}
	aStrSlice := strings.Split(aStr, " ")
	aIntSlice := make([]int, len(aStrSlice))
	for i := range aStrSlice {
		aIntSlice[i], _ = strconv.Atoi(aStrSlice[i])
	}
	sort.Slice(aIntSlice, func(i, j int) bool {
		return aIntSlice[i] > aIntSlice[j]
	})
	sum := 0
	for i := range aIntSlice {
		if i%2 == 0 {
			sum += aIntSlice[i]
		} else {
			sum -= aIntSlice[i]
		}
	}
	fmt.Println(sum)
}
