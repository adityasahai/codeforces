// https://codeforces.com/problemset/problem/1080/B
package main

import (
	"fmt"
	"strconv"
	"bufio"
	"strings"
	"os"
)

func formula(pos int64) int64 {
	if pos % 2 == 0 {
		return int64(pos/2)
	}
	return int64(pos-1)/2 - pos
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numInputs, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	for numInputs > 0 {
		scanner.Scan()
		inputs := strings.Split(scanner.Text(), " ")
		start, _ := strconv.ParseInt(inputs[0], 10, 32)
		last, _ := strconv.ParseInt(inputs[1], 10, 32)
		fmt.Println(int64(formula(last)) - int64(formula(start-1)))
		numInputs--
	}
}