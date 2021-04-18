package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var nStr, aStr string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		nStr = scanner.Text()
	}
	_, _ = strconv.Atoi(nStr)
	if scanner.Scan() {
		aStr = scanner.Text()
	}
	a := strings.Split(aStr, " ")
	count := 0
	for {
		for i := range a {
			aa, _ := strconv.Atoi(a[i])
			if aa%2 == 1 {
				fmt.Println(count)
				return
			}
			a[i] = strconv.Itoa(aa / 2)
		}
		count += 1
	}
}
