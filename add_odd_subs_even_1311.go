// https://codeforces.com/problemset/problem/1311/A
package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func findMoves(x, y string) {
	a, _ := strconv.ParseInt(x, 10, 32)
	b, _ := strconv.ParseInt(y, 10, 32)
	var ans string
	if a == b {
		ans = "0"
	} else if (a % 2) == (b % 2) {
		if a > b {
			ans = "1"
		} else {
			ans = "2"
		}
	} else {
		if a < b {
			ans = "1"
		} else {
			ans = "2"
		}
	}
	fmt.Println(ans)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputs, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	var i int64
	for i = 0; i < inputs; i++ {
		var ab []string
		scanner.Scan()
		ab = strings.Split(scanner.Text(), " ")
		findMoves(ab[0], ab[1])
	}
}