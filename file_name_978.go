// https://codeforces.com/problemset/problem/978/B
package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Scan()
	word := scanner.Text()
	var countx int = 0
	var countrem int = 0
	for _, r := range word {
		if r == 'x' {
			if countx == 2 {
				countrem++
			} else {
				countx++
			}
		} else {
			countx = 0
		}
	}
	fmt.Printf("%d", countrem)
}