// https://codeforces.com/problemset/problem/791/A
package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanf("%d %d\n", &a, &b)
	var years int = 0
	for a <= b {
		years++
		a *= 3
		b *= 2
	}
	fmt.Printf("%d", years)
}
