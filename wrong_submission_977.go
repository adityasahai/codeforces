// https://codeforces.com/problemset/problem/1030/A
package main

import (
	"fmt"
)

func wrongSubstraction() int {
	var number, iter int
	fmt.Scanf("%d %d\n", &number, &iter)

	for i := 0; i < iter; i++ {
		if number%10 == 0 {
			number /= 10
		} else {
			number--
		}
	}
	return number
}

func main() {
	fmt.Printf("%d", wrongSubstraction())
}
