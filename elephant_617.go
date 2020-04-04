// https://codeforces.com/problemset/problem/617/A
package main

import (
	"fmt"
)

var steps = []int{5, 4, 3, 2, 1}

func main() {
	var dest int
	var count int = 0
	fmt.Scanf("%d\n", &dest)

	for dest > 0 {
		for _, step := range steps {
			if step <= dest {
				dest -= step
				count++
				break
			}
		}
	}
	fmt.Printf("%d", count)

}
