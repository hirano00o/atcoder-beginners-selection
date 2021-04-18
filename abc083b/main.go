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
	strs := strings.Split(str, " ")
	nn, aa, bb := strs[0], strs[1], strs[2]
	n, _ := strconv.Atoi(nn)
	a, _ := strconv.Atoi(aa)
	b, _ := strconv.Atoi(bb)

	fSum := 0
	for i := 1; i <= n; i++ {
		s := strconv.Itoa(i)
		sum := 0
		for ss := range s {
			ssInt, _ := strconv.Atoi(string(s[ss]))
			sum += ssInt
		}
		if a <= sum && sum <= b {
			fSum += i
		}
	}
	fmt.Println(fSum)
}
