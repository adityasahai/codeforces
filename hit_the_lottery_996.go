// https://codeforces.com/problemset/problem/996/A
package main

import (
	"fmt"
)

var denominations []int = []int{100, 20, 10, 5, 1}

func main() {
	var dollars, count int
	count = 0
	fmt.Scanf("%d\n", &dollars)
	for _, coin := range denominations {
		for dollars >= coin {
			dollars -= coin
			count++
		}
	}
	fmt.Printf("%d",count)
}